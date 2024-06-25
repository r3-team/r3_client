package send

import (
	"encoding/json"
	"fmt"
	"r3_client/log"
	"r3_client/types"
	"r3_client/websocket/connection"
	"r3_client/websocket/transaction"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var (
	logContext = "websocket"
)

func Do(instanceId uuid.UUID, requests []types.Request) error {

	conn, exists := connection.Get(instanceId)
	if !exists || conn == nil {
		return fmt.Errorf("websocket connection is closed")
	}

	// register transaction (for handling response later)
	trans, err := transaction.Register(requests)
	if err != nil {
		return err
	}

	// send message as JSON
	transJson, err := json.Marshal(trans)
	if err != nil {
		return err
	}
	log.Info(logContext, fmt.Sprintf("sends: %s", transJson))
	return conn.WriteMessage(websocket.TextMessage, transJson)
}
