package tray

import (
	"fmt"
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/file/open"
	"r3_client/icon"
	"r3_client/install"
	"r3_client/log"
	"r3_client/types"

	"fyne.io/systray"
)

type menuItem struct {
	check   bool // is checkbox
	checked bool // checkbox is active
	fn      func()
	title   map[string]string // lang:caption
}

var (
	filesShow  []types.File // last accessed files in order
	logContext = "systray"
	isReady    = false

	// captions
	items = map[string]map[string]string{
		"title": map[string]string{
			"de_de": "REI3 Client",
			"en_us": "REI3 client",
		},
		"conNo": map[string]string{
			"de_de": "[nicht verbunden]",
			"en_us": "[not connected]",
		},
		"conYes": map[string]string{
			"de_de": "[verbunden]",
			"en_us": "[connected]",
		},
		"config": map[string]string{
			"de_de": "Konfigdatei öffnen",
			"en_us": "Open config file",
		},
		"logs": map[string]string{
			"de_de": "Logs öffnen",
			"en_us": "Open logs",
		},
		"startup": map[string]string{
			"de_de": "Autostart",
			"en_us": "Auto start",
		},
		"debug": map[string]string{
			"de_de": "Debug-Modus",
			"en_us": "Debug mode",
		},
		"quit": map[string]string{
			"de_de": "Beenden",
			"en_us": "Quit",
		},
	}

	// menu items
	mTitle   *systray.MenuItem
	mConNo   *systray.MenuItem
	mConYes  *systray.MenuItem
	mFile0   *systray.MenuItem
	mFile1   *systray.MenuItem
	mFile2   *systray.MenuItem
	mConfig  *systray.MenuItem
	mLogs    *systray.MenuItem
	mStartup *systray.MenuItem
	mDebug   *systray.MenuItem
	mQuit    *systray.MenuItem
)

func Fill() {
	// set tray
	lang := config.File.LanguageCode
	systray.SetTitle(items["title"][lang])
	systray.SetIcon(icon.Down)

	// set menu items
	mTitle = systray.AddMenuItem(items["title"][lang], "")
	mTitle.Disable()
	systray.AddSeparator()
	mConNo = systray.AddMenuItem(items["conNo"][lang], "")
	mConNo.Disable()
	mConYes = systray.AddMenuItem(items["conYes"][lang], "")
	mConYes.Disable()
	systray.AddSeparator()
	mFile0 = systray.AddMenuItem("-", "")
	mFile1 = systray.AddMenuItem("-", "")
	mFile2 = systray.AddMenuItem("-", "")
	mFile0.Hide()
	mFile1.Hide()
	mFile2.Hide()
	systray.AddSeparator()
	mConfig = systray.AddMenuItem(items["config"][lang], "")
	mLogs = systray.AddMenuItem(items["logs"][lang], "")
	systray.AddSeparator()
	mStartup = systray.AddMenuItemCheckbox(items["startup"][lang], "", config.File.AutoStart)
	mDebug = systray.AddMenuItemCheckbox(items["debug"][lang], "", config.File.Debug)
	systray.AddSeparator()
	mQuit = systray.AddMenuItem(items["quit"][lang], "")
	SetConnected(false)

	isReady = true

	go func() {
		for {
			select {
			case <-mFile0.ClickedCh:
				openFile(0)
			case <-mFile1.ClickedCh:
				openFile(1)
			case <-mFile2.ClickedCh:
				openFile(2)
			case <-mConfig.ClickedCh:
				open.WithLocalSystem(filepath.Join(config.GetPathApp(), "config.json"), false)
			case <-mLogs.ClickedCh:
				open.WithLocalSystem(filepath.Join(config.GetPathApp(), "client.log"), false)
			case <-mStartup.ClickedCh:
				config.File.AutoStart = !config.File.AutoStart
				if err := config.WriteFile(); err != nil {
					continue
				}
				if err := install.App(); err != nil {
					continue
				}
				if mStartup.Checked() {
					mStartup.Uncheck()
				} else {
					mStartup.Check()
				}
			case <-mDebug.ClickedCh:
				config.File.Debug = !config.File.Debug
				if err := config.WriteFile(); err != nil {
					continue
				}
				log.SetDebug(config.File.Debug)
				if mDebug.Checked() {
					mDebug.Uncheck()
				} else {
					mDebug.Check()
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func openFile(fileIndex int) {
	if fileIndex >= len(filesShow) {
		return
	}
	f := filesShow[fileIndex]
	if err := open.WithLocalSystem(filepath.Join(os.TempDir(), f.DirName, f.FileName), false); err != nil {
		log.Error(logContext, "failed to open file", err)
	}
}

func SetConnected(v bool) {
	if v {
		mConNo.Hide()
		mConYes.Show()
		systray.SetIcon(icon.Neutral)
	} else {
		mConNo.Show()
		mConYes.Hide()
		systray.SetIcon(icon.Down)
	}
}

func SetFiles(files []types.File) {
	log.Info(logContext, fmt.Sprintf("is updating last %d accessed files", len(files)))

	for i, f := range files {
		switch i {
		case 0:
			mFile0.SetTitle(f.FileName)
			mFile0.Show()
		case 1:
			mFile1.SetTitle(f.FileName)
			mFile1.Show()
		case 2:
			mFile2.SetTitle(f.FileName)
			mFile2.Show()
		}
	}
	if len(files) < 3 {
		mFile2.Hide()
	}
	if len(files) < 2 {
		mFile1.Hide()
	}
	if len(files) < 1 {
		mFile0.Hide()
	}
	filesShow = files
}
