package job

import (
	"fmt"
	"time"

	"r3_client/call"
	"r3_client/log"
	"r3_client/websocket"
)

type job struct {
	fn             func() error
	fnAfterSuccess func() error
	intervalSec    int64
	lastRan        int64
	logName        string
}

var (
	running bool

	jobs = []job{job{
		fn:          log.RotateIfNecessary,
		intervalSec: 86400,
		lastRan:     -1,
		logName:     "logRotate",
	}, job{
		fn:          websocket.Connect,
		intervalSec: 5,
		lastRan:     -1,
		logName:     "websocketConnect",
	}, job{
		fn:          call.Authenticate,
		intervalSec: 5,
		lastRan:     -1,
		logName:     "callAuthenticate",
	}}
)

func Start() {
	running = true

	// first start wait time
	time.Sleep(time.Second)

	for {
		if !running {
			break
		}
		now := time.Now().Unix()

		for i, j := range jobs {
			if j.lastRan+j.intervalSec > now {
				continue
			}

			if err := j.fn(); err != nil {
				log.Error("job", fmt.Sprintf("'%s' failed", j.logName), err)
			} else {
				log.Info("job", fmt.Sprintf("'%s' succeeded", j.logName))
			}
			jobs[i].lastRan = now
		}
		time.Sleep(time.Second * 5)
	}
}

func Stop() {
	running = false
}
