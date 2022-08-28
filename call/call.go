package call

import (
	"encoding/json"
	"r3_client/config"
	"r3_client/log"
	"r3_client/types"
	"r3_client/websocket"
)

var logContext = "authentication"

func Authenticate() error {
	for instanceId, inst := range config.GetInstances() {

		if config.GetIsAuthenticated(instanceId) {
			continue
		}

		// prepare request
		payload := types.RequestPayloadLogin{
			LoginId:    inst.LoginId,
			TokenFixed: inst.TokenFixed,
		}
		payloadJson, err := json.Marshal(payload)
		if err != nil {
			log.Error(logContext, "failed marshal request", err)
			continue
		}

		// send transaction
		if err := websocket.Send(instanceId, []types.Request{
			types.Request{
				Ressource: "auth",
				Action:    "tokenFixed",
				Payload:   payloadJson,
			},
		}); err != nil {
			log.Error(logContext, "failed to send websocket request", err)
			continue
		}
	}
	return nil
}
