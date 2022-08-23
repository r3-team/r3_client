package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"r3_client/config"
	"r3_client/file"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/types"

	"github.com/gorilla/websocket"
)

var (
	conn             *websocket.Conn // nil if closed
	logContext       = "websocket"
	transactionNrMap = make(map[uint64]types.RequestTransaction)
	serverUrl        string // URL of websocket server
)

func SetServerUrl(v string) {
	serverUrl = v
}

func Connect() error {
	if conn != nil {
		return nil
	}

	var err error
	conn, _, err = websocket.DefaultDialer.Dial(serverUrl, nil)
	return err
}
func Disconnect() {
	if conn != nil {
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))

		conn.Close()
		conn = nil
		config.SetAuthToken("")
	}
}

func HandleReceived() {
	for {
		if conn == nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Error(logContext, "failed to read message", err)
			Disconnect()
			continue
		}
		log.Info(logContext, fmt.Sprintf("received: %s", message))

		var res types.ResponseTransaction
		if err := json.Unmarshal(message, &res); err != nil {
			log.Error(logContext, "failed to unmarshal message", err)
			continue
		}
		if res.Error != "" {
			log.Error(logContext, "response returned error", errors.New(res.Error))
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

				if err := file.Open(resPayload.AttributeId,
					resPayload.FileId, resPayload.FileHash, resPayload.FileName,
					resPayload.ChooseApp); err != nil {

					log.Error(logContext, "failed to open file", err)
					continue
				}
			}
			continue
		}

		// check if transaction was sent out
		req, exists := transactionNrMap[res.TransactionNr]
		if !exists {
			log.Error(logContext, "response invalid", errors.New("transaction not recognized"))
			continue
		}

		// process authentication messages
		if len(req.Requests) == 1 && req.Requests[0].Ressource == "auth" {
			var resPayload types.ResponsePayloadLogin
			if err := json.Unmarshal(res.Responses[0].Payload, &resPayload); err != nil {
				log.Error(logContext, "failed to unmarshal response payload", err)
				continue
			}
			config.SetAuthToken(resPayload.Token)
			continue
		}

		// process regular messages
		// ... nothing yet
	}
}

func Send(requests []types.Request) error {
	if conn == nil {
		return fmt.Errorf("websocket connection is closed")
	}

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
