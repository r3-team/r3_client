//go:build darwin

package install

import (
	"fmt"
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/tools"
)

const (
	binName    = "r3_client"
	binNameNew = "r3_client_new" // for updater
	binNameOld = "r3_client_old" // for updater
	lnkRef     = "com.lsw.r3_client"
	lnkName    = "com.lsw.r3_client.plist"

	desktopFile = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>Label</key>
    <string>%s</string>
    <key>ProgramArguments</key>
      <array>
        <string>%s</string>
      </array>
    <key>RunAtLoad</key>
    <true/>
    <key>AbandonProcessGroup</key>
    <true/>
  </dict>
</plist>`
)

func getFilePathBin() string {
	return filepath.Join(config.GetPathApp(), binName)
}
func getFilePathCnf() string {
	return filepath.Join(config.GetPathApp(), config.GetFileName())
}
func getFilePathLnk() string {
	return filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents", lnkName)
}
func createLnk(filePathLnk string, filePathBin string) error {

	startupDir := filepath.Join(os.Getenv("HOME"), "Library", "LaunchAgents")
	exists, err := tools.Exists(startupDir)
	if err != nil {
		return err
	}
	if !exists {
		if err := os.Mkdir(startupDir, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(filePathLnk, []byte(fmt.Sprintf(desktopFile, lnkRef, filePathBin)), 0755)
}
