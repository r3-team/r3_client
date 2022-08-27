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
	filesShow     []types.File // last accessed files in order
	logContext    = "systray"
	isConnected   = false
	isLoadingDown = false
	isLoadingUp   = false
	isReady       = false

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
		"ssl": map[string]string{
			"de_de": "SSL verwenden",
			"en_us": "Use SSL",
		},
		"sslVerify": map[string]string{
			"de_de": "SSL verifizieren",
			"en_us": "Verify SSL",
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
	mTitle     *systray.MenuItem
	mConNo     *systray.MenuItem
	mConYes    *systray.MenuItem
	mFile0     *systray.MenuItem
	mFile1     *systray.MenuItem
	mFile2     *systray.MenuItem
	mFile3     *systray.MenuItem
	mFile4     *systray.MenuItem
	mConfig    *systray.MenuItem
	mLogs      *systray.MenuItem
	mStartup   *systray.MenuItem
	mSsl       *systray.MenuItem
	mSslVerify *systray.MenuItem
	mDebug     *systray.MenuItem
	mQuit      *systray.MenuItem
)

func Fill() {
	// set tray
	lang := config.File.LanguageCode
	systray.SetTitle(items["title"][lang])
	systray.SetIcon(icon.Neutral)

	// title
	mTitle = systray.AddMenuItem(items["title"][lang], "")
	mTitle.Disable()

	// connection details
	systray.AddSeparator()
	mConNo = systray.AddMenuItem(items["conNo"][lang], "")
	mConNo.Disable()
	mConNo.Show()
	mConYes = systray.AddMenuItem(items["conYes"][lang], "")
	mConYes.Disable()
	mConYes.Hide()

	// last accessed files
	systray.AddSeparator()
	mFile0 = systray.AddMenuItem("-", "")
	mFile1 = systray.AddMenuItem("-", "")
	mFile2 = systray.AddMenuItem("-", "")
	mFile3 = systray.AddMenuItem("-", "")
	mFile4 = systray.AddMenuItem("-", "")

	// file open actions
	systray.AddSeparator()
	mConfig = systray.AddMenuItem(items["config"][lang], "")
	mLogs = systray.AddMenuItem(items["logs"][lang], "")

	// toggle actions
	systray.AddSeparator()
	mStartup = systray.AddMenuItemCheckbox(items["startup"][lang], "", config.File.AutoStart)
	mSsl = systray.AddMenuItemCheckbox(items["ssl"][lang], "", config.File.Ssl)
	mSslVerify = systray.AddMenuItemCheckbox(items["sslVerify"][lang], "", config.File.SslVerify)
	if !config.File.Ssl {
		mSslVerify.Hide()
	}
	mDebug = systray.AddMenuItemCheckbox(items["debug"][lang], "", config.File.Debug)

	// quite
	systray.AddSeparator()
	mQuit = systray.AddMenuItem(items["quit"][lang], "")

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
			case <-mFile3.ClickedCh:
				openFile(3)
			case <-mFile4.ClickedCh:
				openFile(4)
			case <-mConfig.ClickedCh:
				open.WithLocalSystem(filepath.Join(config.GetPathApp(), config.GetFileName()), false)
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
			case <-mSsl.ClickedCh:
				config.File.Ssl = !config.File.Ssl
				if err := config.WriteFile(); err != nil {
					continue
				}
				if config.File.Ssl {
					mSsl.Check()
					mSslVerify.Show()
				} else {
					mSsl.Uncheck()
					mSslVerify.Hide()
				}
			case <-mSslVerify.ClickedCh:
				config.File.SslVerify = !config.File.SslVerify
				if err := config.WriteFile(); err != nil {
					continue
				}
				if config.File.SslVerify {
					mSslVerify.Check()
				} else {
					mSslVerify.Uncheck()
				}
			case <-mDebug.ClickedCh:
				config.File.Debug = !config.File.Debug
				if err := config.WriteFile(); err != nil {
					continue
				}
				log.SetDebug(config.File.Debug)
				if config.File.Debug {
					mDebug.Check()
				} else {
					mDebug.Uncheck()
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

func SetIcon() {
	// 1st prio: not connected
	if !isConnected {
		systray.SetIcon(icon.Down)
		return
	}

	// 2nd prio: uploading
	if isLoadingUp {
		systray.SetIcon(icon.Upload)
		return
	}

	// 3rd prio: downloading
	if isLoadingDown {
		systray.SetIcon(icon.Download)
		return
	}

	// if nothing else: neutral
	systray.SetIcon(icon.Neutral)
}

func SetLoadingDown(v bool) {
	isLoadingDown = v
	SetIcon()
}
func SetLoadingUp(v bool) {
	isLoadingUp = v
	SetIcon()
}
func SetConnected(v bool) {
	isConnected = v
	if v {
		mConNo.Hide()
		mConYes.Show()
	} else {
		mConNo.Show()
		mConYes.Hide()
	}
	SetIcon()
}

func SetFiles(files []types.File) {
	log.Info(logContext, fmt.Sprintf("is updating last %d accessed files",
		len(files)))

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
		case 3:
			mFile3.SetTitle(f.FileName)
			mFile3.Show()
		case 4:
			mFile4.SetTitle(f.FileName)
			mFile4.Show()
		}
	}
	if len(files) < 5 {
		mFile4.Hide()
	}
	if len(files) < 4 {
		mFile3.Hide()
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
