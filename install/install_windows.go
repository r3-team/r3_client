//go:build windows

package install

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

const (
	binName = "r3_client.exe"
	lnkName = "r3_client.lnk"

	psScript = `$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("%s")
$Shortcut.TargetPath = "%s"
$Shortcut.Save()`
)

func getFilePathBin(appDir string) string {
	return filepath.Join(appDir, binName)
}
func getFilePathLnk(userDir string) string {
	return filepath.Join([]string{userDir, "AppData", "Roaming", "Microsoft",
		"Windows", "Start Menu", "Programs", "Startup", lnkName}...)
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
