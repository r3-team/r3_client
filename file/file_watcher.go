package file

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/file/load"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/tray"
	"r3_client/types"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gofrs/uuid"
)

var (
	watcher                 *fsnotify.Watcher
	watcherChecksPerFile    = 5
	watcherChecksInterval   = time.Millisecond * 2000
	watcherFileIdsActive    = make(map[uuid.UUID]bool)
	watcherFileIdsActive_mx = sync.Mutex{}
)

func WatcherStart() error {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event, isOpen := <-watcher.Events:
				if !isOpen {
					return
				}

				// do not only check for write events
				// MS office uses create/rename operations when writing changes between temp and the original file
				if event.Op != fsnotify.Create && event.Op != fsnotify.Write && event.Op != fsnotify.Rename {
					continue
				}

				// get last directory name
				dirName := filepath.Base(filepath.Dir(event.Name))
				fileName := filepath.Base(event.Name)

				files_mx.Lock()
				for fileId, f := range files {
					if f.DirName == dirName && f.FileName == fileName {

						log.Info(logContext, fmt.Sprintf("recognized file event '%s' (%s) for file ID '%s'",
							event.Name, event.Op.String(), fileId))

						watcherFileIdsActive_mx.Lock()
						if _, exists := watcherFileIdsActive[fileId]; !exists {
							watcherFileIdsActive[fileId] = true
							go watcherReactToWrite(event.Name, fileId, f)
						}
						watcherFileIdsActive_mx.Unlock()
					}
				}
				files_mx.Unlock()
			case err, isOpen := <-watcher.Errors:
				if !isOpen {
					return
				}
				log.Error(logContext, "has encountered an error", err)
			}
		}
	}()
	return nil
}
func WatcherStop() {
	if watcher != nil {
		if err := watcher.Close(); err != nil {
			log.Error(logContext, "failed to close file watcher", err)
		}
		watcher = nil
	}
}

func watcherAdd(path string) error {
	log.Info(logContext, fmt.Sprintf("added watcher on '%s'", path))
	return watcher.Add(path)
}
func watcherRemove(path string) error {
	log.Info(logContext, fmt.Sprintf("removed watcher on '%s'", path))
	return watcher.Remove(path)
}

func watcherReactToWrite(filePath string, fileId uuid.UUID, f types.File) {
	defer watcherReleaseFile(fileId)

	// wait shorty after the event to wait for file locks to be released
	time.Sleep(time.Millisecond * 500)

	var fileInfoLast fs.FileInfo

	for checks := 0; checks < watcherChecksPerFile; checks++ {

		if checks != 0 {
			// on every subsequent event, wait for regular interval to expire
			time.Sleep(watcherChecksInterval)
		}

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Error(logContext, "failed to check file info", err)
			continue
		}

		if checks == 0 {
			// check at least twice to recognize file being written to
			fileInfoLast = fileInfo
			continue
		} else {
			if fileInfoLast.ModTime() != fileInfo.ModTime() || fileInfoLast.Size() != fileInfo.Size() {
				log.Info(logContext, "recognized changed file being written to, ignore change for now")
				fileInfoLast = fileInfo
				continue
			}
		}

		// it seems file is not being written to, check file itself
		// ignore 0 byte files (usual when new file was created and writing has not started)
		if fileInfo.Size() == 0 {
			log.Info(logContext, "recognized empty file, ignore change for now")
			continue
		}

		// check for actual file hash changes
		fileHash, err := tools.GetFileHash(filePath)
		if err != nil {
			log.Error(logContext, "failed to check file hash", err)
			continue
		}
		if f.FileHash == fileHash {
			log.Info(logContext, "recognized identical file hash, ignore change for now")
			continue
		}

		// upload file
		log.Info(logContext, "recognized file change, trigger file version upload")

		inst, err := config.GetInstance(f.InstanceId)
		if err != nil {
			log.Error(logContext, "failed to upload new file version", fmt.Errorf("unknown instance"))
			break
		}

		scheme := "https"
		if !config.GetSsl() {
			scheme = "http"
		}
		url := fmt.Sprintf("%s://%s:%d/data/upload", scheme, inst.HostName, inst.HostPort)
		params := map[string]string{
			"token":       config.GetAuthToken(f.InstanceId),
			"attributeId": f.AttributeId.String(),
			"fileId":      fileId.String(),
		}

		if err := upload(url, params, f.FileName, filePath); err != nil {
			log.Error(logContext, "failed to upload new file version", err)
			break
		}

		log.Info(logContext, "uploaded new file version successfully")

		// upload successful, update file hash
		setFileHash(fileId, fileHash)
		break
	}

}
func watcherReleaseFile(fileId uuid.UUID) {
	watcherFileIdsActive_mx.Lock()
	delete(watcherFileIdsActive, fileId)
	watcherFileIdsActive_mx.Unlock()
}

func upload(url string, params map[string]string, fileName string, filePath string) error {
	tray.SetLoadingUp(true)
	defer tray.SetLoadingUp(false)

	return load.Up(url, params, fileName, filePath)
}
