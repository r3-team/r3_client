package keyboard_listen

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"r3_client/config"
	"r3_client/log"
	"r3_client/types"
	"r3_client/websocket"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/gofrs/uuid"
	hook "github.com/robotn/gohook"
)

var (
	logContext string = "hotkey"
)

func Start() {
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

				// execute callback function if used
				if action.JsFunctionId.Valid {
					if err := executeJsFunction(instanceId, action.JsFunctionArgs, action.JsFunctionId.UUID); err != nil {
						log.Error(logContext, "failed to execute action on hotkey", err)
					}
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

func executeJsFunction(instanceId uuid.UUID, args []string, jsFunctionId uuid.UUID) error {
	argValues := make([]interface{}, 0)

	for _, arg := range args {
		var err error
		var value interface{}
		switch arg {
		case "clipboard":
			value, err = clipboard.ReadAll()
			if err != nil {
				return err
			}
		case "hostname":
			value, err = os.Hostname()
			if err != nil {
				return err
			}
		case "username":
			user, err := user.Current()
			if err != nil {
				return err
			}
			value = user.Username
		case "windowTitle":
			value = robotgo.GetTitle()
		default:
			value = nil
		}
		argValues = append(argValues, value)
	}

	// prepare request
	payload := types.RequestPayloadJsFunctionCall{
		JsFunctionId: jsFunctionId,
		Arguments:    argValues,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// send request to browser session
	return websocket.Send(instanceId, []types.Request{
		{
			Ressource: "device",
			Action:    "browserCallJsFunction",
			Payload:   payloadJson,
		},
	})
}
