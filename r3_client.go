package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"r3_client/config"
	"r3_client/event"
	"r3_client/file"
	"r3_client/install"
	"r3_client/job"
	"r3_client/lock"
	"r3_client/log"
	"r3_client/tools"
	"r3_client/tray"
	"r3_client/ws"

	"fyne.io/systray"
)

var (
	logContext = "system"

	// overwritten by build parameters
	appVersion = "0.1.2.3"
)

func main() {
	config.SetAppVersion(appVersion)

	// listen to global shutdown channel
	signal.Notify(config.OsExit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	go func() {
		<-config.OsExit
		systray.Quit()
	}()

	// start systray
	systray.Run(onReady, onExit)
}

func quitWithErr(message string, err error) {
	log.Error(logContext, message, err)
	systray.Quit()
}

func onReady() {

	// get user home dir
	userDir, err := os.UserHomeDir()
	if err != nil {
		quitWithErr("failed to ascertain user home directory", err)
		return
	}

	// define application user directory (create if required)
	var appDir string
	switch runtime.GOOS {
	case "darwin", "linux":
		appDir = filepath.Join(userDir, ".r3")
	case "windows":
		appDir = filepath.Join(userDir, "AppData", "Local", "r3")
	default:
		quitWithErr("failed to start", fmt.Errorf("unsupported runtime environment '%s'", runtime.GOOS))
		return
	}

	exists, err := tools.Exists(appDir)
	if err != nil {
		quitWithErr("failed to check application user directory", err)
		return
	}
	if !exists {
		if err := os.Mkdir(appDir, 0755); err != nil {
			quitWithErr("failed to create application user directory", err)
			return
		}
	}

	// define paths
	config.SetPathApp(appDir)
	config.SetPathUser(userDir)
	file.SetFilePathCache(filepath.Join(appDir, config.GetFileNameCache()))
	log.SetFilePath(filepath.Join(appDir, config.GetFileNameLog()))

	// install application, errors should not block execution
	if err := install.App(); err != nil {
		log.Error(logContext, "failed to install application", err)
	}

	// check whether another instance of the application is running
	if err := lock.GetExclusive(); err != nil {
		quitWithErr("failed to get exclusive access to lock file", err)
		return
	}

	// load config file
	if err := config.ReadFile(); err != nil {
		quitWithErr("failed to read config file", err)
		return
	}

	// apply logging settings from config file
	log.SetDebug(config.GetDebug())

	// set auto start, errors should not block execution
	if err := install.AutoStart(); err != nil {
		log.Error(logContext, "failed to setup auto start", err)
	}

	// fill system tray
	tray.SetDefaults()

	// start file system watcher
	if err := file.WatcherStart(); err != nil {
		quitWithErr("failed to start file system watcher", err)
		return
	}

	// restore handled files from cache
	if err := file.CacheRestore(); err != nil {
		quitWithErr("failed to restore file cache", err)
		return
	}

	// start regular jobs
	go job.Start()
}

func onExit() {
	for instanceId := range config.GetInstances() {
		event.ExecuteOn(instanceId, "onDisconnect")
	}

	lock.Release()
	job.Stop()
	file.WatcherStop()
	ws.DisconnectAll()
}
