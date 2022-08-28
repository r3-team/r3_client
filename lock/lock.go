package lock

import (
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/tools"
)

var (
	file *os.File // application lock file handler
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
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0400)
	return err
}

func Release() {
	file.Close()
}
