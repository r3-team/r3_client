package action

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"r3_client/types"
	"r3_client/ws/ws_connect"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/gofrs/uuid"
)

func Do(instanceId uuid.UUID, clientEvent types.Event) error {

	// block invalid executions
	if clientEvent.Action == "callJsFunction" && !clientEvent.JsFunctionId.Valid {
		return fmt.Errorf("no ID given for JS function call")
	}
	if clientEvent.Action == "callPgFunction" && !clientEvent.PgFunctionId.Valid {
		return fmt.Errorf("no ID given for PG function call")
	}

	args := make([]interface{}, 0)

	for _, arg := range clientEvent.Arguments {
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
		args = append(args, value)
	}

	// send execution request to the server
	payload := types.RequestPayloadClientEventExec{
		Id:        clientEvent.Id,
		Arguments: args,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return ws_connect.SendToInstance(instanceId, []types.Request{{
		Ressource: "clientEvent",
		Action:    "exec",
		Payload:   payloadJson,
	}})
}
