package install

import (
	"fmt"
	"os"
	"path/filepath"

	"r3_client/config"
	"r3_client/log"
	"r3_client/tools"
)

var (
	logContext = "installer"
)

func App() error {

	// get OS dependent paths
	filePathBinNow, err := os.Executable()
	if err != nil {
		return err
	}
	filePathBin := getFilePathBin()
	filePathCnf := getFilePathCnf()
	filePathLnk := getFilePathLnk()

	if filePathBinNow != filePathBin {
		// app is started outside of its directory

		// install app to application directory if not there already
		exists, err := tools.Exists(filePathBin)
		if err != nil {
			return err
		}
		if !exists {
			if err := copyApp(filePathBinNow, filePathBin); err != nil {
				return err
			}

		} else {
			// overwrite app if hash changed
			fileHashBinNow, err := tools.GetFileHash(filePathBinNow)
			if err != nil {
				return err
			}
			fileHashBin, err := tools.GetFileHash(filePathBin)
			if err != nil {
				return err
			}

			if fileHashBinNow != fileHashBin {
				if err := copyApp(filePathBinNow, filePathBin); err != nil {
					return err
				}
			}
		}

		// copy config file, if within the same directory
		currDir := filepath.Dir(filePathBinNow)
		filePathCnfCurrDir := filepath.Join(currDir, config.GetFileName())

		exists, err = tools.Exists(filePathCnfCurrDir)
		if err != nil {
			return err
		}
		if exists {
			if err := tools.FileCopy(filePathCnfCurrDir, filePathCnf, false); err != nil {
				return err
			}
			if err := config.LoadCreateFile(); err != nil {
				return err
			}
		}
	}

	// install/deinstall auto start
	exists, err := tools.Exists(filePathLnk)
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

func copyApp(pathSrc string, pathDst string) error {
	log.Info(logContext, fmt.Sprintf("is copying executable to application directory '%s'",
		pathDst))

	if err := tools.FileCopy(pathSrc, pathDst, false); err != nil {
		return err
	}
	return os.Chmod(pathDst, 0744)
}
