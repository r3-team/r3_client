package file

import (
	"encoding/json"
	"os"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/types"
)

var (
	filesCachePath = "" // cache of handled files, for recovery after app restart
)

func SetFilePathCache(v string) {
	filesCachePath = v
}

func CacheStore() {
	files_mx.Lock()
	cache := types.FilesSaved{
		Files: files,
	}
	files_mx.Unlock()

	cacheJson, err := json.MarshalIndent(cache, "", "\t")
	if err != nil {
		log.Error(logContext, "failed to update file cache", err)
		return
	}
	if err := os.WriteFile(filesCachePath, cacheJson, 0644); err != nil {
		log.Error(logContext, "failed to update file cache JSON file", err)
		return
	}
}

func CacheRestore() error {
	exists, err := tools.Exists(filesCachePath)
	if err != nil {
		return err
	}
	if !exists {
		updateTray()
		return nil
	}

	cacheJson, err := tools.GetFileContents(filesCachePath, true)
	if err != nil {
		return err
	}

	var cache types.FilesSaved
	if err := json.Unmarshal(cacheJson, &cache); err != nil {
		return err
	}

	files_mx.Lock()
	files = cache.Files
	for _, f := range files {
		if err := watcherAdd(GetDirPath(f.DirName)); err != nil {
			log.Error(logContext, "failed to add file path to watcher", err)
			continue
		}
	}
	files_mx.Unlock()

	updateTray()
	return nil
}
