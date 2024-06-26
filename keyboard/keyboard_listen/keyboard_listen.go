package keyboard_listen

import (
	"fmt"
	"r3_client/event/action"
	"r3_client/log"
	"r3_client/types"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gofrs/uuid"
	hook "github.com/robotn/gohook"
)

var (
	logContext string = "hotkey"
	running    atomic.Bool
)

func Start(instanceIdMapEvents map[uuid.UUID][]types.Event) {

	if running.Load() {
		hook.End()
		time.Sleep(time.Millisecond * 100)
	}

	running.Store(true)
	defer running.Store(false)

	// collect all hotkey events for all instances
	for instanceId, events := range instanceIdMapEvents {

		for _, event := range events {
			if event.Event != "onHotkey" {
				continue
			}

			keys := make([]string, 0)

			if event.HotkeyModifier1 != "" {
				keys = append(keys, strings.ToLower(event.HotkeyModifier1))
			}
			if event.HotkeyModifier2 != "" {
				keys = append(keys, strings.ToLower(event.HotkeyModifier2))
			}
			keys = append(keys, strings.ToLower(event.HotkeyChar))

			hook.Register(hook.KeyDown, keys, func(e hook.Event) {
				log.Info(logContext, fmt.Sprintf("reacting to hotkey %s\n", keys))

				if err := action.Do(instanceId, event); err != nil {
					log.Error(logContext, "failed to execute client event action 'onHotkey'", err)
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
