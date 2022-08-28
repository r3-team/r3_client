package file

import (
	"fmt"
	"os"
	"r3_client/config"
	"r3_client/log"
	"r3_client/tools"

	"github.com/gofrs/uuid"
)

// attempt to delete outdated files
func CleanupFiles() error {

	now := tools.GetTimeUnix()
	idsDelete := make([]uuid.UUID, 0)
	secKeepFor := config.GetKeepFilesSec()

	files_mx.Lock()
	for id, f := range files {
		if f.Touched+secKeepFor < now {
			idsDelete = append(idsDelete, id)
		}
	}

	for _, id := range idsDelete {
		dirPath := GetDirPath(files[id].DirName)
		filePath := GetFilePath(files[id].DirName, files[id].FileName)

		log.Info(logContext, fmt.Sprintf("cleaning up file '%s'", filePath))
		if err := os.Remove(filePath); err != nil {
			log.Error(logContext, "failed to cleanup file", err)
		}

		if err := watcherRemove(dirPath); err != nil {
			log.Error(logContext, "failed to remove file system watcher", err)
		}

		log.Info(logContext, fmt.Sprintf("cleaning up file directory '%s'", dirPath))
		if err := os.RemoveAll(dirPath); err != nil {
			log.Error(logContext, "failed to cleanup file directory", err)
		}
		delete(files, id)
	}
	files_mx.Unlock()

	if len(idsDelete) != 0 {
		CacheStore()
		updateTray()
	}
	return nil
}
