package file

import (
	"fmt"
	"path/filepath"
	"r3_client/config"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/types"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/gofrs/uuid"
)

var (
	watcher                 *fsnotify.Watcher
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

				if event.Op == fsnotify.Write {

					// get last directory name
					dirName := filepath.Base(filepath.Dir(event.Name))
					fileName := filepath.Base(event.Name)

					files_mx.Lock()
					for fileId, f := range files {
						if f.DirName == dirName && f.FileName == fileName {

							watcherFileIdsActive_mx.Lock()
							if _, exists := watcherFileIdsActive[fileId]; !exists {
								watcherFileIdsActive[fileId] = true
								go watcherReactToWrite(event.Name, fileId, f)
							}
							watcherFileIdsActive_mx.Unlock()
						}
					}
					files_mx.Unlock()
				}
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
	if err := watcher.Close(); err != nil {
		log.Error(logContext, "failed to close file watcher", err)
	}
	watcher = nil
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

	fileHash, err := tools.GetFileHash(filePath)
	if err != nil {
		log.Error(logContext, "failed to check file hash", err)
		return
	}

	if f.FileHash == fileHash {
		return
	}
	log.Info(logContext, "recognized file change, trigger update")

	scheme := "https"
	if !config.File.Ssl {
		scheme = "http"
	}
	url := fmt.Sprintf("%s://%s:%d/data/upload", scheme,
		config.File.HostName, config.File.HostPort)

	params := map[string]string{
		"token":       config.GetAuthToken(),
		"attributeId": f.AttributeId.String(),
		"fileId":      fileId.String(),
	}

	if err := upload(url, params, f.FileName, filePath); err != nil {
		log.Error(logContext, "failed to upload new file version", err)
		return
	}
	log.Info(logContext, "uploaded new file version successfully")

	// upload successful, update file hash
	files_mx.Lock()
	if f, exists := files[fileId]; exists {
		f.FileHash = fileHash
		files[fileId] = f
	}
	files_mx.Unlock()
}
func watcherReleaseFile(fileId uuid.UUID) {
	watcherFileIdsActive_mx.Lock()
	if _, exists := watcherFileIdsActive[fileId]; exists {
		delete(watcherFileIdsActive, fileId)
	}
	watcherFileIdsActive_mx.Unlock()
}
