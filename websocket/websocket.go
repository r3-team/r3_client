package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"r3_client/config"
	"r3_client/event"
	"r3_client/file"
	"r3_client/keyboard/keyboard_listen"
	"r3_client/keyboard/keyboard_type"
	"r3_client/log"
	"r3_client/tray"
	"r3_client/types"
	"r3_client/websocket/connection"
	"r3_client/websocket/send"
	"r3_client/websocket/transaction"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var (
	read_mx    sync.RWMutex
	logContext = "websocket"

	// requests
	reqClientAppGetBuild = types.Request{
		Ressource: "clientApp",
		Action:    "getBuild",
		Payload:   nil,
	}
	reqClientEventGet = types.Request{
		Ressource: "clientEvent",
		Action:    "get",
		Payload:   nil,
	}
)

// connect to all known instances via websocket
func Connect() error {
	for instanceId, instance := range config.GetInstances() {

		conn, err := connection.Connect(instanceId, instance)
		if err != nil {
			log.Error(logContext, fmt.Sprintf("failed to connect to instance '%s'", instanceId), err)
			continue
		}

		go handleReceived(instanceId, conn)
	}
	return nil
}

func handleReceived(instanceId uuid.UUID, conn *websocket.Conn) {
	for {
		read_mx.Lock()
		_, message, err := conn.ReadMessage()
		read_mx.Unlock()

		if err != nil {
			// connection closed, abort
			if conn == nil {
				return
			}

			// connection error, close
			log.Error(logContext, "encountered read error, closing connection", err)

			config.SetInstanceToken(instanceId, "")
			tray.SetConnected(instanceId, false)
			connection.Disconnect(instanceId, conn, true)
			return
		}
		log.Info(logContext, fmt.Sprintf("received: %s", message))

		var res types.ResponseTransaction
		if err := json.Unmarshal(message, &res); err != nil {
			log.Error(logContext, "failed to unmarshal message", err)
			continue
		}
		if res.Error != "" {
			log.Error(logContext, "response returned error", errors.New(res.Error))
			tray.SetConnected(instanceId, false)
			continue
		}

		// handle unrequested messages
		if res.TransactionNr == 0 && len(res.Responses) == 1 {

			var resUnreq types.UnreqResponseTransaction
			if err := json.Unmarshal(message, &resUnreq); err != nil {
				log.Error(logContext, "failed to unmarshal unrequested transaction", err)
				continue
			}

			switch resUnreq.Responses[0].Ressource {
			case "clientEventsChanged", "reauthorized":
				if err := send.Do(instanceId, []types.Request{reqClientEventGet}); err != nil {
					log.Error(logContext, "failed to send websocket request", err)
				}
			case "fileRequested":
				var resPayload types.UnreqResponsePayloadFileRequested
				if err := json.Unmarshal(resUnreq.Responses[0].Payload, &resPayload); err != nil {
					log.Error(logContext, "failed to unmarshal unrequested response payload", err)
					continue
				}

				if err := file.Open(instanceId, resPayload.AttributeId,
					resPayload.FileId, resPayload.FileHash, resPayload.FileName,
					resPayload.ChooseApp); err != nil {

					log.Error(logContext, "failed to open file", err)
					continue
				}
			case "keystrokesRequested":
				var resPayload string
				if err := json.Unmarshal(resUnreq.Responses[0].Payload, &resPayload); err != nil {
					log.Error(logContext, "failed to unmarshal unrequested response payload", err)
					continue
				}
				keyboard_type.Do(resPayload)
			}
			continue
		}

		// check if transaction was sent out
		trans, exists := transaction.Get(res.TransactionNr)
		if !exists {
			log.Error(logContext, "response invalid", errors.New("transaction not recognized"))
			continue
		}
		transaction.Deregister(trans.TransactionNr)

		// process authentication response
		if isAuthTransaction(trans) {
			var resPayload types.ResponsePayloadLogin
			if err := json.Unmarshal(res.Responses[0].Payload, &resPayload); err != nil {
				log.Error(logContext, "failed to unmarshal response payload", err)
				continue
			}
			config.SetInstanceToken(instanceId, resPayload.Token)

			// get app build & client events after successful authentication
			if err := send.Do(instanceId, []types.Request{reqClientAppGetBuild, reqClientEventGet}); err != nil {
				log.Error(logContext, "failed to send websocket request", err)
			}
			continue
		}

		// process regular responses
		for i, req := range trans.Requests {
			switch req.Ressource {
			case "clientApp":
				switch req.Action {
				case "getBuild":
					var build int
					if err := json.Unmarshal(res.Responses[i].Payload, &build); err != nil {
						log.Error(logContext, "failed to unmarshal response payload", err)
						continue
					}
					if build > config.GetAppVersionBuild() {
						tray.SetUpdateAvailable(instanceId)
					}
				}
			case "clientEvent":
				switch req.Action {
				case "get":
					var resPayload struct {
						ClientEvents          []types.Event                  `json:"clientEvents"`
						ClientEventIdMapLogin map[uuid.UUID]types.EventLogin `json:"clientEventIdMapLogin"`
					}
					if err := json.Unmarshal(res.Responses[i].Payload, &resPayload); err != nil {
						log.Error(logContext, "failed to unmarshal response payload", err)
						continue
					}

					// parse client events
					clientEvents := make([]types.Event, 0)
					for _, ce := range resPayload.ClientEvents {
						if ce.Event != "onHotkey" {
							clientEvents = append(clientEvents, ce)
							continue
						}

						// hotkey events must be registered for the current login, apply hotkey overwrites
						lce, exists := resPayload.ClientEventIdMapLogin[ce.Id]
						if exists {
							ce.HotkeyChar = lce.HotkeyChar
							ce.HotkeyModifier1 = lce.HotkeyModifier1
							ce.HotkeyModifier2 = lce.HotkeyModifier2
							clientEvents = append(clientEvents, ce)
						}
					}

					event.SetByInstanceId(instanceId, clientEvents)
					event.ExecuteOn(instanceId, "onConnect")

					// refresh keyboard listeners on change
					go keyboard_listen.Start()
				}
			}
		}
	}
}

func isAuthTransaction(t types.RequestTransaction) bool {
	return len(t.Requests) == 1 && t.Requests[0].Ressource == "auth"
}
