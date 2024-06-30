//go:build windows

package install

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"r3_client/config"
	"syscall"
)

const (
	binName    = "r3_client.exe"
	binNameNew = "r3_client_new.exe" // for updater
	binNameOld = "r3_client_old.exe" // for updater
	lnkName    = "r3_client.lnk"

	psScript = `$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("%s")
$Shortcut.TargetPath = "%s"
$Shortcut.Save()`
)

func getFilePathBin() string {
	return filepath.Join(config.GetPathApp(), binName)
}
func getFilePathCnf() string {
	return filepath.Join(config.GetPathApp(), config.GetFileName())
}
func getFilePathLnk() string {
	return filepath.Join([]string{config.GetPathUser(), "AppData", "Roaming",
		"Microsoft", "Windows", "Start Menu", "Programs", "Startup", lnkName}...)
}
func createLnk(filePathLnk string, filePathBin string) error {
	psCmd, err := exec.LookPath("powershell.exe")
	if err != nil {
		return err
	}
	cmd := exec.Command(psCmd, []string{
		"-NoProfile",
		"-NonInteractive",
		fmt.Sprintf(psScript, filePathLnk, filePathBin),
	}...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Run()
}
