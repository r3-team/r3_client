package event

import (
	"fmt"
	"r3_client/event/action"
	"r3_client/keyboard/keyboard_listen"
	"r3_client/log"
	"r3_client/types"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

var (
	access_mx           = &sync.RWMutex{}
	instanceIdMapEvents = make(map[uuid.UUID][]types.Event)
	logContext          = "event"
)

func GetByInstanceId(instanceId uuid.UUID) []types.Event {
	access_mx.RLock()
	defer access_mx.RUnlock()

	ce, ok := instanceIdMapEvents[instanceId]
	if !ok {
		return []types.Event{}
	}
	return ce
}

func SetByInstanceId(instanceId uuid.UUID, events []types.Event) {
	access_mx.Lock()
	instanceIdMapEvents[instanceId] = events
	access_mx.Unlock()

	// refresh keyboard listeners on change
	go keyboard_listen.Start(instanceIdMapEvents)
}

func ExecuteEvents(instanceId uuid.UUID, eventTarget string) {
	access_mx.RLock()
	defer access_mx.RUnlock()

	events, exists := instanceIdMapEvents[instanceId]
	if !exists {
		return
	}

	anyCall := false

	for _, ev := range events {
		if ev.Event == eventTarget && ev.Action == "callJsFunction" && ev.JsFunctionId.Valid {
			if err := action.CallFunction(instanceId, ev.JsFunctionArgs, ev.JsFunctionId.UUID); err != nil {
				log.Error(logContext, fmt.Sprintf("failed to execute %s", eventTarget), err)
			}
			anyCall = true
		}
	}

	// add short sleep after sending onDisconnect call
	if eventTarget == "onDisconnect" && anyCall {
		time.Sleep(time.Second * 1)
	}
}
