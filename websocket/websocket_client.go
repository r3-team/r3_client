package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"r3_client/config"
	"r3_client/file"
	"r3_client/keyboard/keyboard_type"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/tray"
	"r3_client/types"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var (
	logContext           = "websocket"
	instanceIdMapConn    = make(map[uuid.UUID]*websocket.Conn) // map of websocket connections, key: instance ID
	instanceIdMapConn_mx = sync.Mutex{}
	transactionNrMap     = make(map[uint64]types.RequestTransaction)
	transactionNrMap_mx  = sync.Mutex{}
)

// connect to all known instances via websocket
func Connect() error {
	for instanceId, inst := range config.GetInstances() {

		instanceIdMapConn_mx.Lock()
		conn, connExists := instanceIdMapConn[instanceId]
		instanceIdMapConn_mx.Unlock()

		if connExists && conn != nil {
			continue
		}

		scheme := "wss"
		if !config.GetSsl() {
			scheme = "ws"
		}

		header := http.Header{}
		header.Add("User-Agent", "r3-client-fat")

		tlsConfig := config.GetTlsConfig()
		dialer := websocket.Dialer{
			TLSClientConfig: &tlsConfig,
		}

		var err error
		conn, _, err = dialer.Dial(fmt.Sprintf("%s://%s:%d/websocket", scheme, inst.HostName, inst.HostPort), header)

		// update system tray
		tray.SetConnected(instanceId, err == nil)

		if err != nil {
			log.Warning(logContext, "failed to connect", err)
			continue
		}

		instanceIdMapConn_mx.Lock()
		instanceIdMapConn[instanceId] = conn
		instanceIdMapConn_mx.Unlock()
		go handleReceived(instanceId, conn)
	}
	return nil
}

func DisconnectAll() {
	log.Info(logContext, "is closing all connections")

	instanceIdMapConn_mx.Lock()
	defer instanceIdMapConn_mx.Unlock()

	for instanceId, _ := range config.GetInstances() {

		conn, exists := instanceIdMapConn[instanceId]
		if !exists || conn == nil {
			continue
		}

		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

		conn.Close()
	}
}

func handleReceived(instanceId uuid.UUID, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			// connection closed, abort
			if conn == nil {
				return
			}

			// connection error, close
			log.Error(logContext, "encountered read error, closing connection", err)

			config.SetInstanceToken(instanceId, "")
			tray.SetConnected(instanceId, false)

			instanceIdMapConn_mx.Lock()
			conn.Close()
			delete(instanceIdMapConn, instanceId)
			instanceIdMapConn_mx.Unlock()
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
			case "keystrokesDo":
				var resPayload types.UnreqResponsePayloadKeystrokesDo
				if err := json.Unmarshal(resUnreq.Responses[0].Payload, &resPayload); err != nil {
					log.Error(logContext, "failed to unmarshal unrequested response payload", err)
					continue
				}
				keyboard_type.Do(resPayload.Strokes)
			}
			continue
		}

		// check if transaction was sent out
		transactionNrMap_mx.Lock()
		trans, exists := transactionNrMap[res.TransactionNr]
		transactionNrMap_mx.Unlock()

		if !exists {
			log.Error(logContext, "response invalid", errors.New("transaction not recognized"))
			continue
		}

		// process authentication messages
		if isAuthTransaction(trans) {
			var resPayload types.ResponsePayloadLogin
			if err := json.Unmarshal(res.Responses[0].Payload, &resPayload); err != nil {
				log.Error(logContext, "failed to unmarshal response payload", err)
				continue
			}
			config.SetInstanceToken(instanceId, resPayload.Token)
			continue
		}

		// process regular messages
		// ... nothing yet
	}
}

func Send(instanceId uuid.UUID, requests []types.Request) error {

	conn, exists := instanceIdMapConn[instanceId]
	if !exists || conn == nil {
		return fmt.Errorf("websocket connection is closed")
	}

	transactionNrMap_mx.Lock()
	defer transactionNrMap_mx.Unlock()

	// create transaction
	trans := types.RequestTransaction{
		Requests: requests,
	}

	// register transaction (for handling response later)
	var number uint64
	for true {
		number = uint64(tools.RandNumber(100000, 499999))

		if _, exists := transactionNrMap[number]; !exists {
			trans.TransactionNr = number
			transactionNrMap[number] = trans
			break
		}
	}

	// send message as JSON
	transJson, err := json.Marshal(trans)
	if err != nil {
		return err
	}
	log.Info(logContext, fmt.Sprintf("sends: %s", transJson))
	return conn.WriteMessage(websocket.TextMessage, transJson)
}

func isAuthTransaction(t types.RequestTransaction) bool {
	return len(t.Requests) == 1 && t.Requests[0].Ressource == "auth"
}
