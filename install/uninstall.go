package install

import (
	"os"
	"r3_client/config"
	"r3_client/log"
	"r3_client/tools"
)

func Remove() error {
	// disable auto start
	config.SetAutoStart(false)
	if err := AutoStart(); err != nil {
		// do not abort removal if autostart could not be disabled
		log.Error(logContext, "failed to remove autostart", err)
	}

	// delete the configuration file, making the client application unusable
	filePathCnf := getFilePathCnf()
	exists, err := tools.Exists(filePathCnf)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return os.Remove(filePathCnf)
}
