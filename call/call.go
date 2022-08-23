package call

import (
	"encoding/json"
	"r3_client/config"
	"r3_client/types"
	"r3_client/websocket"
)

func Authenticate() error {

	if config.GetIsAuthenticated() {
		return nil
	}

	// prepare request
	payload := types.RequestPayloadLogin{
		LoginId:    config.File.LoginId,
		TokenFixed: config.File.TokenFixed,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// send transaction
	return websocket.Send([]types.Request{
		types.Request{
			Ressource: "auth",
			Action:    "tokenFixed",
			Payload:   payloadJson,
		},
	})
}
