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
func SetLoadingDown(v bool) {
	isLoadingDown = v
	updateIcon()
}
func SetLoadingUp(v bool) {
	isLoadingUp = v
	updateIcon()
}
func SetFiles(files []types.File) {
	access_mx.Lock()
	filesShow = files
	access_mx.Unlock()
	FillMenu()
}
