//go:build windows

package install

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"r3_client/config"
)

const (
	binName = "r3_client.exe"
	lnkName = "r3_client.lnk"

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
	return exec.Command(psCmd, []string{
		"-NoProfile",
		"-NonInteractive",
		fmt.Sprintf(psScript, filePathLnk, filePathBin),
	}...).Run()
}
