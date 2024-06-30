package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"r3_client/config"
	"r3_client/log"
	"r3_client/tools"

	"github.com/go-vgo/robotgo"
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

	if filepath.Dir(filePathBinNow) == filepath.Dir(filePathBin) {
		// app runs inside its directory, everything is fine
		return nil
	}

	// app is started outside of its directory

	// kill currently running client
	pidCurr := os.Getpid()
	pidsAll, err := robotgo.FindIds("r3_client")
	if err == nil {
		for _, pid := range pidsAll {
			if pid != pidCurr {
				log.Info(logContext, fmt.Sprintf("found running client process (pid %d), killing it", pid))

				p, err := os.FindProcess(pid)
				if err != nil {
					log.Error(logContext, "failed to query existing client process", err)
				}
				if err := p.Kill(); err != nil {
					log.Error(logContext, "failed to kill existing client process", err)
				}
			}
		}
	}

	// install app to its directory if not there already or outdated
	exists, err := tools.Exists(filePathBin)
	if err != nil {
		return err
	}
	if !exists {
		if err := copyApp(filePathBinNow, filePathBin); err != nil {
			return err
		}
	} else {
		// compare old to new client hash
		fileHashBinNow, err := tools.GetFileHash(filePathBinNow)
		if err != nil {
			return err
		}
		fileHashBin, err := tools.GetFileHash(filePathBin)
		if err != nil {
			return err
		}

		if fileHashBinNow != fileHashBin {
			log.Info(logContext, "found outdated client version, overwriting it")
			if err := copyApp(filePathBinNow, filePathBin); err != nil {
				return err
			}
		}
	}

	// copy config file, if none is set
	exists, err = tools.Exists(filePathCnf)
	if err != nil {
		return err
	}
	if !exists {
		// check: directory of binary and user download directory (Windows/MacOS at least)
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
			// we only need to copy the config file once
			break
		}
	}

	// start app in target dir and then exit
	log.Info(logContext, fmt.Sprintf("starting client in target directory '%s' and exiting", filePathBin))
	cmd := exec.Command(filePathBin)
	if err := cmd.Start(); err != nil {
		return err
	}
	config.OsExit <- syscall.SIGTERM

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
