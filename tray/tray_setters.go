package tray

import (
	"r3_client/icon"
	"r3_client/types"

	"fyne.io/systray"
	"github.com/gofrs/uuid"
)

func updateIcon() {
	access_mx.Lock()
	defer access_mx.Unlock()

	// 1st prio: any instance not connected
	for _, connected := range instanceIdMapConnected {
		if !connected {
			systray.SetIcon(icon.Down)
			return
		}
	}

	// 2nd prio: uploading
	if isLoadingUp {
		systray.SetIcon(icon.Upload)
		return
	}

	// 3rd prio: downloading
	if isLoadingDown {
		systray.SetIcon(icon.Download)
		return
	}

	// if nothing else: neutral
	systray.SetIcon(icon.Neutral)
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
