package tray

import (
	"r3_client/types"
	"time"

	"github.com/gofrs/uuid"
)

func SetActionDone(v bool) {
	isActionDone.Store(v)
	updateIcon()

	if v {
		go func() {
			// state "action done" is just shown temporarily
			time.Sleep(time.Second * 3)
			isActionDone.Store(false)
			updateIcon()
		}()
	}
}
func SetConnected(instanceId uuid.UUID, v bool) {
	access_mx.Lock()
	connected, exists := instanceIdMapConnected[instanceId]
	instanceIdMapConnected[instanceId] = v
	access_mx.Unlock()

	// update system tray, if connection state changed
	if !exists || connected != v {
		updateIcon()
		FillMenu()
	}
}
func SetFiles(files []types.File) {
	access_mx.Lock()
	filesShow = files
	access_mx.Unlock()
	FillMenu()
}
func SetLoadingDown(v bool) {
	isLoadingDown.Store(v)
	updateIcon()
}
func SetLoadingUp(v bool) {
	isLoadingUp.Store(v)
	updateIcon()
}
func SetUpdateAvailable(instanceId uuid.UUID) {
	access_mx.Lock()
	isUpdateAvailable = true
	isUpdateAvailableInstanceId = instanceId
	access_mx.Unlock()
	FillMenu()
}
