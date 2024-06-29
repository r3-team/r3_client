package tray

import (
	"r3_client/types"

	"github.com/gofrs/uuid"
)

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
	access_mx.Lock()
	isLoadingDown = v
	access_mx.Unlock()
	updateIcon()
}
func SetLoadingUp(v bool) {
	access_mx.Lock()
	isLoadingUp = v
	access_mx.Unlock()
	updateIcon()
}
func SetUpdateAvailable(instanceId uuid.UUID) {
	access_mx.Lock()
	isUpdateAvailable = true
	isUpdateAvailableInstanceId = instanceId
	access_mx.Unlock()
	FillMenu()
}
