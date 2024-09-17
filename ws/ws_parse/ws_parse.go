package ws_parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"r3_client/config"
	"r3_client/event"
	"r3_client/file"
	"r3_client/keyboard/keyboard_listen"
	"r3_client/keyboard/keyboard_type"
	"r3_client/log"
	"r3_client/tray"
	"r3_client/types"
	"r3_client/ws/ws_trans"

	"github.com/gofrs/uuid"
)

var (
	logContext = "websocket"

	// requests
	reqClientAppGetBuild = types.Request{
		Ressource: "clientApp",
		Action:    "getBuild",
		Payload:   nil,
	}
	reqClientEventGet = types.Request{
		Ressource: "clientEvent",
		Action:    "get",
		Payload:   nil,
	}
)

func Do(instanceId uuid.UUID, message []byte, chanWrite chan []types.Request) error {
	log.Info(logContext, fmt.Sprintf("received: %s", message))

	sendWhenFree := func(requests []types.Request) {
		// write channel is blocked until read is done
		go func() {
			chanWrite <- requests
		}()
	}

	var res types.ResponseTransaction
	if err := json.Unmarshal(message, &res); err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}

	// handle unrequested messages
	if res.TransactionNr == 0 && len(res.Responses) == 1 {

		var resUnreq types.UnreqResponseTransaction
		if err := json.Unmarshal(message, &resUnreq); err != nil {
			return err
		}

		switch resUnreq.Responses[0].Ressource {
		case "clientEventsChanged", "reauthorized":
			sendWhenFree([]types.Request{reqClientEventGet})
		case "fileRequested":
			var resPayload types.UnreqResponsePayloadFileRequested
			if err := json.Unmarshal(resUnreq.Responses[0].Payload, &resPayload); err != nil {
				return err
			}

			if err := file.Open(instanceId, resPayload.AttributeId,
				resPayload.FileId, resPayload.FileHash, resPayload.FileName,
				resPayload.ChooseApp); err != nil {

				return err
			}
		case "keystrokesRequested":
			var resPayload string
			if err := json.Unmarshal(resUnreq.Responses[0].Payload, &resPayload); err != nil {
				return err
			}
			keyboard_type.Do(resPayload)
		}
		return nil
	}

	// check if transaction was sent out
	trans, exists := ws_trans.Get(res.TransactionNr)
	if !exists {
		return errors.New("transaction not recognized")
	}
	ws_trans.Deregister(trans.TransactionNr)

	// process authentication response
	if isAuthTransaction(trans) {
		var resPayload types.ResponsePayloadLogin
		if err := json.Unmarshal(res.Responses[0].Payload, &resPayload); err != nil {
			return err
		}
		config.SetInstanceToken(instanceId, resPayload.Token)

		// get app build & client events after successful authentication
		sendWhenFree([]types.Request{reqClientAppGetBuild, reqClientEventGet})
		return nil
	}

	// process regular responses
	for i, req := range trans.Requests {
		switch req.Ressource {
		case "clientApp":
			switch req.Action {
			case "getBuild":
				var build int
				if err := json.Unmarshal(res.Responses[i].Payload, &build); err != nil {
					log.Error(logContext, "failed to unmarshal response payload", err)
					continue
				}
				if build > config.GetAppVersionBuild() {
					tray.SetUpdateAvailable(instanceId)
				}
			}
		case "clientEvent":
			switch req.Action {
			case "get":
				var resPayload struct {
					ClientEvents          []types.Event                  `json:"clientEvents"`
					ClientEventIdMapLogin map[uuid.UUID]types.EventLogin `json:"clientEventIdMapLogin"`
				}
				if err := json.Unmarshal(res.Responses[i].Payload, &resPayload); err != nil {
					log.Error(logContext, "failed to unmarshal response payload", err)
					continue
				}

				// parse client events
				clientEvents := make([]types.Event, 0)
				for _, ce := range resPayload.ClientEvents {
					if ce.Event != "onHotkey" {
						clientEvents = append(clientEvents, ce)
						continue
					}

					// hotkey events must be registered for the current login, apply hotkey overwrites
					lce, exists := resPayload.ClientEventIdMapLogin[ce.Id]
					if exists {
						ce.HotkeyChar = lce.HotkeyChar
						ce.HotkeyModifier1 = lce.HotkeyModifier1
						ce.HotkeyModifier2 = lce.HotkeyModifier2
						clientEvents = append(clientEvents, ce)
					}
				}
				event.SetByInstanceId(instanceId, clientEvents)
				event.ExecuteOn(instanceId, "onConnect")

				// refresh keyboard listeners on change
				go keyboard_listen.Start()
			}
		}
	}
	return nil
}

func isAuthTransaction(t types.RequestTransaction) bool {
	return len(t.Requests) == 1 && t.Requests[0].Ressource == "auth"
}
