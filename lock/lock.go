package lock

import (
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/tools"
	"sync"
)

var (
	file    *os.File // application lock file handler
	file_mx = sync.Mutex{}
)

// try to get exclusive access to the application lock file
func GetExclusive() error {
	filePath := filepath.Join(config.GetPathApp(), "r3_client.lock")

	exists, err := tools.Exists(filePath)
	if err != nil {
		return err
	}

	if exists {
		if err := os.Remove(filePath); err != nil {
			// existing lock file cannot be deleted, another instance might be running
			return err
		}
	}
	file_mx.Lock()
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0400)
	file_mx.Unlock()
	return err
}

func Release() {
	file_mx.Lock()
	defer file_mx.Unlock()
	file.Close()
}
