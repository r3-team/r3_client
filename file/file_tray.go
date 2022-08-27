package file

import (
	"r3_client/tray"
	"r3_client/types"
	"sort"

	"github.com/gofrs/uuid"
)

func updateTray() {
	files_mx.Lock()
	defer files_mx.Unlock()

	// show latest 3 touched files in systray
	keys := make([]uuid.UUID, 0, len(files))
	for k := range files {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return files[keys[i]].Touched > files[keys[j]].Touched
	})

	filesShow := make([]types.File, 0)
	for i := 0; i < trayFileCnt && i < len(keys); i++ {
		filesShow = append(filesShow, files[keys[i]])
	}
	tray.SetFiles(filesShow)
}
