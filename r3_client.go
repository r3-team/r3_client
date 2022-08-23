package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"

	"r3_client/config"
	"r3_client/file"
	"r3_client/install"
	"r3_client/job"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/websocket"
)

var logContext = "system"

func main() {

	// get user home dir
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Error(logContext, "failed to ascertain user home directory", err)
		return
	}

	// define application user directory (create if required)
	var appDir string
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "linux":
		appDir = filepath.Join(userDir, ".r3")
	case "windows":
		appDir = filepath.Join(userDir, "AppData", "Local", "r3")
	default:
		log.Error(logContext, "failed to start", fmt.Errorf("unsupported runtime environment '%s'", runtime.GOOS))
		return
	}
	log.Info(logContext, fmt.Sprintf("set application user directory to '%s'", appDir))

	exists, err := tools.Exists(appDir)
	if err != nil {
		log.Error(logContext, "failed check application user directory", err)
		return
	}
	if !exists {
		if err := os.Mkdir(appDir, 0755); err != nil {
			log.Error(logContext, "failed to create application user directory", err)
			return
		}
	}

	// define paths
	config.SetFilePath(filepath.Join(appDir, "config.json"))
	file.SetFilePathCache(filepath.Join(appDir, "files.json"))
	log.SetFilePath(filepath.Join(appDir, "client.log"))

	// load or create config file
	if err := config.LoadCreateFile(); err != nil {
		log.Error(logContext, "failed to load/create config file", err)
		return
	}

	// apply logging settings from config file
	log.SetLevel(config.File.LogLevel)

	// install application
	if err := install.Do(userDir, appDir, config.File.AutoStart); err != nil {
		log.Error(logContext, "failed to install application", err)
		return
	}

	// prepare websocket client
	wsScheme := "wss"
	if !config.File.Ssl {
		wsScheme = "ws"
	}
	websocket.SetServerUrl(fmt.Sprintf("%s://%s:%d/websocket",
		wsScheme, config.File.HostName, config.File.HostPort))

	go websocket.HandleReceived()
	defer websocket.Disconnect()

	// start file system watcher
	if err := file.WatcherStart(); err != nil {
		log.Error(logContext, "failed to start file system watcher", err)
		return
	}
	defer file.WatcherStop()

	// restore handled files from cache
	if err := file.CacheRestore(); err != nil {
		log.Error(logContext, "failed to restore file cache", err)
		return
	}

	// start regular jobs
	go job.Start()
	defer job.Stop()

	// block until interrupted
	osExit := make(chan os.Signal)
	signal.Notify(osExit, os.Interrupt)
	for {
		select {
		case <-osExit:
		}
	}
}
