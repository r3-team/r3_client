package connection

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
	logContext           = "websocket"
	instanceIdMapConn    = make(map[uuid.UUID]*websocket.Conn) // map of websocket connections, key: instance ID
	instanceIdMapConn_mx = sync.Mutex{}
)

func Get(instanceId uuid.UUID) (*websocket.Conn, bool) {
	instanceIdMapConn_mx.Lock()
	defer instanceIdMapConn_mx.Unlock()
	conn, exists := instanceIdMapConn[instanceId]
	return conn, exists
}
func Connect(instanceId uuid.UUID, instance types.Instance) (*websocket.Conn, error) {
	instanceIdMapConn_mx.Lock()
	defer instanceIdMapConn_mx.Unlock()

	conn, exists := instanceIdMapConn[instanceId]
	if exists && conn != nil {
		return conn, nil
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

	var err error
	conn, _, err = dialer.Dial(fmt.Sprintf("%s://%s:%d/websocket",
		scheme, instance.HostName, instance.HostPort), header)

	// update system tray
	tray.SetConnected(instanceId, err == nil)

	if err != nil {
		return nil, err
	}
	instanceIdMapConn[instanceId] = conn
	return conn, nil
}

func Disconnect(instanceId uuid.UUID, conn *websocket.Conn, onError bool) {
	log.Info(logContext, fmt.Sprintf("is closing connection to instance '%s'", instanceId))

	instanceIdMapConn_mx.Lock()
	defer instanceIdMapConn_mx.Unlock()

	if conn != nil {
		if !onError {
			conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		}
		conn.Close()
	}
	delete(instanceIdMapConn, instanceId)
}
func DisconnectAll() {
	for instanceId, _ := range config.GetInstances() {
		conn, exists := instanceIdMapConn[instanceId]
		if exists {
			Disconnect(instanceId, conn, false)
		}
	}
}
