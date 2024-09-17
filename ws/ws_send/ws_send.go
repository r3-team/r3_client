package ws_send

import (
	"encoding/json"
	"fmt"
	"r3_client/log"
	"r3_client/types"
	"r3_client/ws/ws_trans"

	"github.com/gorilla/websocket"
)

var (
	logContext = "websocket"
)

// send a websocket message to a known instance
func Do(ws *websocket.Conn, requests []types.Request) error {

	// register transaction (for handling response later)
	trans, err := ws_trans.Register(requests)
	if err != nil {
		return err
	}

	// send message as JSON
	transJson, err := json.Marshal(trans)
	if err != nil {
		return err
	}
	log.Info(logContext, fmt.Sprintf("sends: %s", transJson))

	return ws.WriteMessage(websocket.TextMessage, transJson)
}

func DoClose(ws *websocket.Conn) error {
	return ws.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
