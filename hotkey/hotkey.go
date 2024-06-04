package hotkey

import (
	"encoding/json"
	"fmt"
	"r3_client/config"
	"r3_client/log"
	"r3_client/types"
	"r3_client/websocket"

	"github.com/go-vgo/robotgo"
	"github.com/gofrs/uuid"
	hook "github.com/robotn/gohook"
)

var (
	logContext string = "hotkey"
)

func LoadAndListen() {
	instances := config.GetInstances()

	// collect all hotkey actions for all instances
	for instanceId, instance := range instances {
		for _, action := range instance.Actions {
			if !action.Hotkey.Active {
				continue
			}

			keys := make([]string, 0)

			if action.Hotkey.Modifier1 != "" {
				keys = append(keys, action.Hotkey.Modifier1)
			}
			if action.Hotkey.Modifier2 != "" {
				keys = append(keys, action.Hotkey.Modifier2)
			}
			keys = append(keys, action.Hotkey.Char)

			hook.Register(hook.KeyDown, keys, func(e hook.Event) {
				log.Info(logContext, fmt.Sprintf("reacting to hotkey %s\n", keys))
				if err := executeJsFunction(instanceId, action.Action, action.JsFunctionId); err != nil {
					log.Error(logContext, "failed to execute action on hotkey", err)
				}
			})
		}
	}

	s := hook.Start()
	<-hook.Process(s)
}

func Stop() {
	hook.End()
}

func executeJsFunction(instanceId uuid.UUID, action string, jsFunctionId uuid.UUID) error {
	args := make([]interface{}, 0)

	switch action {
	case "sendCurrentWindowTitle":
		args = append(args, robotgo.GetTitle())
	}
	log.Info(logContext, fmt.Sprintf("executing action '%s' on instance '%s', args %s\n", action, instanceId, args))

	// prepare request
	payload := types.RequestPayloadJsFunctionCall{
		JsFunctionId: jsFunctionId,
		Arguments:    args,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := websocket.Send(instanceId, []types.Request{
		{
			Ressource: "fatClient",
			Action:    "jsFunctionCalled",
			Payload:   payloadJson,
		},
	}); err != nil {
		return err
	}

	return nil
}
