package install

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"r3_client/config"
	"r3_client/file/load"
	"r3_client/log"
	"r3_client/tools"
	"runtime"
	"syscall"

	"github.com/gofrs/uuid"
)

func Update(instanceId uuid.UUID) error {

	token := config.GetAuthToken(instanceId)
	if token == "" {
		return fmt.Errorf("empty token")
	}

	os := ""
	switch runtime.GOOS {
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			os = "amd64_mac"
		default:
			return fmt.Errorf("unsupported GOARCH")
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			os = "amd64_linux"
		case "arm64":
			os = "arm64_linux"
		default:
			return fmt.Errorf("unsupported GOARCH")
		}
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			os = "amd64_windows"
		default:
			return fmt.Errorf("unsupported GOARCH")
		}
	default:
		return fmt.Errorf("unsupported GOOS")
	}

	// download new version (do not rename prod client yet as download can fail for many reasons)
	inst, err := config.GetInstance(instanceId)
	if err != nil {
		return err
	}

	scheme := "https"
	if !config.GetSsl() {
		scheme = "http"
	}

	// client file paths
	fileClientProd := filepath.Join(config.GetPathApp(), binName)
	fileClientNew := filepath.Join(config.GetPathApp(), binNameNew)
	fileClientOld := filepath.Join(config.GetPathApp(), binNameOld)

	log.Info(logContext, fmt.Sprintf("downloading update to '%s'", fileClientNew))

	url := fmt.Sprintf("%s://%s:%d/client/download/?token=%s&os=%s", scheme, inst.HostName, inst.HostPort, token, os)
	if err := load.Down(url, fileClientNew); err != nil {
		return err
	}

	// rename running client binary to old version
	if err := tools.FileMove(fileClientProd, fileClientOld, false); err != nil {
		return err
	}

	// rename new client binary to production version
	if err := tools.FileMove(fileClientNew, fileClientProd, false); err != nil {
		return err
	}

	// start new client, then exit
	log.Info(logContext, "starting updated client and exiting")

	cmd := exec.Command(fileClientProd)
	if err := cmd.Start(); err != nil {
		return err
	}
	config.OsExit <- syscall.SIGTERM

	return nil
}
