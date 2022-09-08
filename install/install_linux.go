//go:build linux

package install

import (
	"fmt"
	"os"
	"path/filepath"
	"r3_client/config"
)

const (
	binName = "r3_client.bin"
	lnkName = "r3_client.desktop"

	desktopFile = `[Desktop Entry]
Type=Application
Name=REI3 client
Exec=%s
StartupNotify=false
Terminal=false`
)

func getFilePathBin() string {
	return filepath.Join(config.GetPathApp(), binName)
}
func getFilePathCnf() string {
	return filepath.Join(config.GetPathApp(), config.GetFileName())
}
func getFilePathLnk() string {
	var dir string

	if os.Getenv("XDG_CONFIG_HOME") != "" {
		dir = os.Getenv("XDG_CONFIG_HOME")
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return filepath.Join(dir, "autostart", lnkName)
}
func createLnk(filePathLnk string, filePathBin string) error {
	return os.WriteFile(filePathLnk, []byte(fmt.Sprintf(desktopFile, filePathBin)), 0755)
}
