package file

import (
	"encoding/json"
	"os"
	"path/filepath"
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

func CacheStore() error {
	files_mx.Lock()
	cache := types.FilesSaved{
		Files: files,
	}
	files_mx.Unlock()

	cacheJson, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	return os.WriteFile(filesCachePath, cacheJson, 0644)
}

func CacheRestore() error {
	exists, err := tools.Exists(filesCachePath)
	if err != nil {
		return err
	}
	if !exists {
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
	defer files_mx.Unlock()
	for _, f := range files {
		if err := watcherAdd(filepath.Join(tempDir, f.DirName)); err != nil {
			log.Error(logContext, "failed to add file path to watcher", err)
			continue
		}
	}
	return nil
}
