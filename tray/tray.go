package tray

import (
	"path/filepath"
	"r3_client/config"
	"r3_client/file"
	"r3_client/icon"
	"r3_client/install"
	"r3_client/log"

	"fyne.io/systray"
)

type menuItem struct {
	check   bool // is checkbox
	checked bool // checkbox is active
	fn      func()
	title   map[string]string // lang:caption
}

var (
	logContext = "systray"
	items      = map[string]map[string]string{
		"title": map[string]string{
			"de_de": "REI3 Client",
			"en_us": "REI3 client",
		},
		"conNo": map[string]string{
			"de_de": "[Nicht verbunden]",
			"en_us": "[Not connected]",
		},
		"conYes": map[string]string{
			"de_de": "[Verbunden]",
			"en_us": "[Connected]",
		},
		"startup": map[string]string{
			"de_de": "Autostart",
			"en_us": "Auto start",
		},
		"logs": map[string]string{
			"de_de": "Logs öffnen",
			"en_us": "Open logs",
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
	mStartup *systray.MenuItem
	mLogs    *systray.MenuItem
	mDebug   *systray.MenuItem
	mQuit    *systray.MenuItem
)

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
	mLogs = systray.AddMenuItem(items["logs"][lang], "")
	systray.AddSeparator()
	mStartup = systray.AddMenuItemCheckbox(items["startup"][lang], "", config.File.AutoStart)
	mDebug = systray.AddMenuItemCheckbox(items["debug"][lang], "", config.File.Debug)
	systray.AddSeparator()
	mQuit = systray.AddMenuItem(items["quit"][lang], "")
	SetConnected(false)

	go func() {
		for {
			select {
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
			case <-mLogs.ClickedCh:
				file.OpenWithLocalSystem(filepath.Join(config.GetPathApp(), "client.log"), false)
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}
