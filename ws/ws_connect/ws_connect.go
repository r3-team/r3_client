package ws_connect

import (
	"fmt"
	"net/http"
	"r3_client/config"
	"r3_client/log"
	"r3_client/tray"
	"r3_client/types"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var (
	instanceIdMapCon    = make(map[uuid.UUID]types.WsCon) // map of websocket connections, key: instance ID
	instanceIdMapCon_mx = sync.Mutex{}
	logContext          = "websocket"
)

// connect to one instance
// returns websocket connection and whether its a new one
func Connect(instanceId uuid.UUID, instance types.Instance) (types.WsCon, bool, error) {
	instanceIdMapCon_mx.Lock()
	defer instanceIdMapCon_mx.Unlock()

	if con, exists := instanceIdMapCon[instanceId]; exists {
		return con, false, nil
	}

	scheme := "wss"
	if !config.GetSsl() {
		scheme = "ws"
	}

	header := http.Header{}
	header.Add("User-Agent", "r3-client-fat")

	tlsConfig := config.GetTlsConfig()
	dialer := websocket.Dialer{
		TLSClientConfig: &tlsConfig,
	}

	log.Info(logContext, fmt.Sprintf("is connecting to instance '%s'", instanceId))

	ws, _, err := dialer.Dial(fmt.Sprintf("%s://%s:%d/websocket",
		scheme, instance.HostName, instance.HostPort), header)

	// update system tray
	tray.SetConnected(instanceId, err == nil)

	if err != nil {
		return types.WsCon{}, false, err
	}
	con := types.WsCon{
		ChanClose:  make(chan error),
		ChanRead:   make(chan []byte),
		ChanWrite:  make(chan []types.Request),
		InstanceId: instanceId,
		Ws:         ws,
	}

	instanceIdMapCon[instanceId] = con
	log.Info(logContext, fmt.Sprintf("successfully connected to '%s' (open connections: %d)", instanceId, len(instanceIdMapCon)))

	return con, true, nil
}

func Disconnect(instanceId uuid.UUID) {
	instanceIdMapCon_mx.Lock()
	con, exists := instanceIdMapCon[instanceId]
	instanceIdMapCon_mx.Unlock()

	if exists {
		con.ChanClose <- nil
	}
}

func Remove(instanceId uuid.UUID) {
	instanceIdMapCon_mx.Lock()
	delete(instanceIdMapCon, instanceId)
	instanceIdMapCon_mx.Unlock()
}

func SendToInstance(instanceId uuid.UUID, requests []types.Request) error {
	instanceIdMapCon_mx.Lock()
	defer instanceIdMapCon_mx.Unlock()

	con, exists := instanceIdMapCon[instanceId]
	if !exists {
		return fmt.Errorf("websocket connection is not active")
	}
	con.ChanWrite <- requests
	return nil
}
