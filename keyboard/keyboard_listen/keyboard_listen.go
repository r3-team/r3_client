package keyboard_listen

import (
	"fmt"
	"r3_client/event"
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
	starting   atomic.Bool
	running    atomic.Bool
)

// required to call Execute() with different event ID
func getFncForEvent(keys []string, instanceId uuid.UUID, ev types.Event) func(hook.Event) {
	return func(e hook.Event) {
		log.Info(logContext, fmt.Sprintf("reacting to hotkey %s\n", keys))
		action.Do(instanceId, ev)
	}
}

func Start() {
	if starting.Load() {
		return
	}
	starting.Store(true)
	StopIfRunning()

	log.Info(logContext, "registering new global hooks")

	// collect hotkey events for known instances
	for instanceId, events := range event.Get() {

		for _, ev := range events {
			if ev.Event != "onHotkey" {
				continue
			}

			if ev.HotkeyChar == "" {
				log.Error(logContext, "failed to register hotkey", fmt.Errorf("character key is empty"))
				continue
			}

			keys := []string{
				strings.ToLower(ev.HotkeyChar),
				strings.ToLower(ev.HotkeyModifier1),
			}
			if ev.HotkeyModifier2.Valid {
				keys = append(keys, strings.ToLower(ev.HotkeyModifier2.String))
			}
			hook.Register(hook.KeyDown, keys, getFncForEvent(keys, instanceId, ev))
		}
	}

	log.Info(logContext, "starting global hooks")

	s := hook.Start()

	starting.Store(false)
	running.Store(true)
	defer running.Store(false)

	<-hook.Process(s)
}

func StopIfRunning() {
	if running.Load() {
		log.Info(logContext, "stopping global hooks")
		hook.End()
		time.Sleep(time.Millisecond * 500)
	}
}
