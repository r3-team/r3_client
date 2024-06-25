package action

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"r3_client/types"
	"r3_client/websocket/send"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/gofrs/uuid"
)

func CallFunction(instanceId uuid.UUID, args []string, jsFunctionId uuid.UUID) error {
	argValues := make([]interface{}, 0)

	for _, arg := range args {
		var err error
		var value interface{}
		switch arg {
		case "clipboard":
			if clipboard.Unsupported {
				return fmt.Errorf("clipboard access is not supported")
			} else {
				value, err = clipboard.ReadAll()
				if err != nil {
					if strings.Contains(err.Error(), "Element not found") {
						// expected error if clipboard is empty, at least on Windows
						value = nil
					} else {
						return err
					}
				}
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
			return fmt.Errorf("unknown function parameter")
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
	return send.Do(instanceId, []types.Request{{
		Ressource: "device",
		Action:    "browserCallJsFunction",
		Payload:   payloadJson,
	}})
}
