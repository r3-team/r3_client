package ws

import (
	"encoding/json"
	"fmt"
	"r3_client/config"
	"r3_client/log"
	"r3_client/tray"
	"r3_client/types"
	"r3_client/ws/ws_connect"
	"r3_client/ws/ws_parse"
	"r3_client/ws/ws_send"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var (
	logContext = "websocket"
)

func ConnectAll() error {
	for instanceId, instance := range config.GetInstances() {
		con, isNew, err := ws_connect.Connect(instanceId, instance)
		if err != nil {
			log.Error(logContext, fmt.Sprintf("failed to connect to instance '%s'", instanceId), err)
			continue
		}

		if isNew {
			go process(con)
			go listen(con)
		}

		// auth can fail, repeat if not successfully authenticated
		if !config.GetIsAuthenticated(instanceId) {
			authenticate(instanceId, con.ChanWrite)
		}
	}
	return nil
}

func DisconnectAll() {
	for instanceId := range config.GetInstances() {
		ws_connect.Disconnect(instanceId)
	}
}

func listen(con types.WsCon) {
	for {
		_, message, err := con.Ws.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Error(logContext, "failed to read websocket message", err)
			}
			con.ChanClose <- err
			return
		} else {
			con.ChanRead <- message
		}
	}
}

func process(con types.WsCon) {
	for {
		select {
		case message := <-con.ChanRead:
			if err := ws_parse.Do(con.InstanceId, message, con.ChanWrite); err != nil {
				log.Error(logContext, "failed to process websocket message", err)
			}
		case requests := <-con.ChanWrite:
			if err := ws_send.Do(con.Ws, requests); err != nil {
				log.Error(logContext, "failed to send websocket request", err)
			}
		case errClose := <-con.ChanClose:
			if errClose == nil {
				log.Info(logContext, fmt.Sprintf("closing connection '%s' with close message", con.InstanceId))
				if err := ws_send.DoClose(con.Ws); err != nil {
					log.Error(logContext, "failed to send close message", err)
				}
				if err := con.Ws.Close(); err != nil {
					log.Error(logContext, "failed to close websocket connection", err)
				}
			}

			config.SetInstanceToken(con.InstanceId, "")
			tray.SetConnected(con.InstanceId, false)
			ws_connect.Remove(con.InstanceId)
			return
		}
	}
}

func authenticate(instanceId uuid.UUID, chanWrite chan []types.Request) {

	log.Info(logContext, fmt.Sprintf("started authenticating against instance '%s'", instanceId))
	inst, err := config.GetInstance(instanceId)
	if err != nil {
		log.Error(logContext, "failed to authenticate", err)
		return
	}

	payload := types.RequestPayloadLogin{
		LoginId:    inst.LoginId,
		TokenFixed: inst.TokenFixed,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Error(logContext, "failed to authenticate", err)
		return
	}

	chanWrite <- []types.Request{{
		Ressource: "auth",
		Action:    "tokenFixed",
		Payload:   payloadJson,
	}}
}
