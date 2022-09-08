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

	if filepath.Dir(filePathBinNow) != filepath.Dir(filePathBin) {
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

		// copy config file, if none is set yet
		exists, err = tools.Exists(filePathCnf)
		if err != nil {
			return err
		}
		if exists {
			return nil
		}

		// check: Directory of binary and user download directory (Windows/MacOS at least)
		paths := []string{
			filepath.Join(filepath.Dir(filePathBinNow), config.GetFileName()),
			filepath.Join(config.GetPathUser(), "Downloads", config.GetFileName()),
		}
		for _, path := range paths {
			exists, err = tools.Exists(path)
			if err != nil {
				return err
			}
			if !exists {
				continue
			}
			if err := tools.FileCopy(path, filePathCnf, false); err != nil {
				return err
			}
			if err := config.ReadFile(); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func AutoStart() error {
	filePathBin := getFilePathBin()
	filePathLnk := getFilePathLnk()

	exists, err := tools.Exists(filePathLnk)
	if err != nil {
		return err
	}

	if config.GetAutoStart() && !exists {
		// copy link to binary to startup folder
		log.Info(logContext, fmt.Sprintf("is setting auto start link file at '%s'",
			filePathLnk))

		return createLnk(filePathLnk, filePathBin)

	} else if !config.GetAutoStart() && exists {
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
