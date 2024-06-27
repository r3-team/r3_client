package event

import (
	"fmt"
	"r3_client/event/action"
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

func Get() map[uuid.UUID][]types.Event {
	access_mx.RLock()
	defer access_mx.RUnlock()
	return instanceIdMapEvents
}

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
}

// execute actions for specific event target (onConnect, onDisconnect, onHotkey)
func ExecuteOn(instanceId uuid.UUID, eventTarget string) {
	access_mx.RLock()
	events, exists := instanceIdMapEvents[instanceId]
	access_mx.RUnlock()

	if !exists {
		return
	}

	anyCall := false

	for _, ev := range events {
		if ev.Event == eventTarget {
			if err := action.Do(instanceId, ev); err != nil {
				log.Error(logContext, fmt.Sprintf("failed to execute client event actions for '%s'", eventTarget), err)
			}
			anyCall = true
		}
	}

	// add short sleep after sending onDisconnect call
	if eventTarget == "onDisconnect" && anyCall {
		time.Sleep(time.Second * 1)
	}
}
