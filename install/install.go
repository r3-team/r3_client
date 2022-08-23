package install

import (
	"fmt"
	"os"

	"r3_client/config"
	"r3_client/log"
	"r3_client/tools"
)

var (
	logContext = "installer"
)

func App() error {

	// get OS dependent paths
	filePathBin := getFilePathBin(config.GetPathApp())
	filePathLnk := getFilePathLnk(config.GetPathUser())

	// install app to user application directory if not there already
	exists, err := tools.Exists(filePathBin)
	if err != nil {
		return err
	}
	if !exists {
		log.Info(logContext, fmt.Sprintf("is copying executable to application directory '%s'",
			filePathBin))

		filePathBinNow, err := os.Executable()
		if err != nil {
			return err
		}
		if err := tools.FileCopy(filePathBinNow, filePathBin, false); err != nil {
			return err
		}
	}

	// install/deinstall auto start
	exists, err = tools.Exists(filePathLnk)
	if err != nil {
		return err
	}

	if config.File.AutoStart && !exists {
		// copy link to binary to startup folder
		log.Info(logContext, fmt.Sprintf("is setting auto start link file at '%s'",
			filePathLnk))

		return createLnk(filePathLnk, filePathBin)
	}
	if !config.File.AutoStart && exists {
		// remove link to binary from startup folder
		log.Info(logContext, fmt.Sprintf("is removing auto start link file at '%s'",
			filePathLnk))

		return os.Remove(filePathLnk)
	}
	return nil

}
