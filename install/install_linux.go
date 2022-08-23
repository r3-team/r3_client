//go:build linux

package install

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	binName = "r3_client"
	lnkName = "r3_client.deskop"

	desktopFile = `[Desktop Entry]
Type=Application
Name=REI3 client
Exec=%s
StartupNotify=false
Terminal=false`
)

func getFilePathBin(appDir string) string {
	return filepath.Join(appDir, binName)
}
func getFilePathLnk(userDir string) string {
	var dir string

	if os.Getenv("XDG_CONFIG_HOME") != "" {
		dir = os.Getenv("XDG_CONFIG_HOME")
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return filepath.Join(dir, "autostart", lnkName)
}
func createLnk(filePathLnk string, filePathBin string) error {
	return os.WriteFile(filePathLnk, []byte(fmt.Sprintf(desktopFile, filePathBin)), 0644)
}