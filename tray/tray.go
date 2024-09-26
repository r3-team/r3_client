package tray

import (
	"fmt"
	"os"
	"path/filepath"
	"r3_client/config"
	"r3_client/file/open"
	"r3_client/icon"
	"r3_client/icon/dark"
	"r3_client/install"
	"r3_client/log"
	"r3_client/types"
	"sync"
	"sync/atomic"
	"syscall"

	"fyne.io/systray"
	"github.com/gofrs/uuid"
)

var (
	access_mx = sync.Mutex{}

	filesShow              = make([]types.File, 0)    // last accessed files in order
	instanceIdMapConnected = make(map[uuid.UUID]bool) // map of instances that are connected, key: instance ID
	logContext             = "systray"
	title                  = "" // system tray title

	isActionDone  atomic.Bool
	isLoadingDown atomic.Bool
	isLoadingUp   atomic.Bool

	isFilledOnce                = false
	isUpdateAvailable           = false
	isUpdateAvailableInstanceId uuid.UUID

	// captions
	items = map[string]map[string]string{
		"title": {
			"de_de": "REI3 Client",
			"en_us": "REI3 client",
		},
		"conNo": {
			"de_de": "[nicht verbunden]",
			"en_us": "[not connected]",
		},
		"conYes": {
			"de_de": "[verbunden]",
			"en_us": "[connected]",
		},
		"config": {
			"de_de": "Konfigdatei öffnen",
			"en_us": "Open config file",
		},
		"logs": {
			"de_de": "Logs öffnen",
			"en_us": "Open logs",
		},
		"darkIcon": {
			"de_de": "Dunkles Icon",
			"en_us": "Dark icon",
		},
		"startup": {
			"de_de": "Autostart",
			"en_us": "Auto start",
		},
		"ssl": {
			"de_de": "SSL verwenden",
			"en_us": "Use SSL",
		},
		"sslVerify": {
			"de_de": "SSL verifizieren",
			"en_us": "Verify SSL",
		},
		"debug": {
			"de_de": "Debug-Logging",
			"en_us": "Debug logging",
		},
		"quit": {
			"de_de": "Beenden",
			"en_us": "Quit",
		},
		"uninstall0": {
			"de_de": "Deinstallieren",
			"en_us": "Uninstall",
		},
		"uninstall1": {
			"de_de": "REI3 Client deinstallieren?",
			"en_us": "Uninstall REI3 client?",
		},
		"update": {
			"de_de": "Update installieren",
			"en_us": "Install update",
		},
	}

	// menu items
	itemsCleared = make(chan bool)
	mTitle       *systray.MenuItem
	mFile0       *systray.MenuItem
	mFile1       *systray.MenuItem
	mFile2       *systray.MenuItem
	mFile3       *systray.MenuItem
	mFile4       *systray.MenuItem
	mConfig      *systray.MenuItem
	mLogs        *systray.MenuItem
	mDarkIcon    *systray.MenuItem
	mStartup     *systray.MenuItem
	mSsl         *systray.MenuItem
	mSslVerify   *systray.MenuItem
	mDebug       *systray.MenuItem
	mUpdate      *systray.MenuItem
	mUninstall0  *systray.MenuItem
	mUninstall1  *systray.MenuItem
	mQuit        *systray.MenuItem
)

func SetDefaults() {
	access_mx.Lock()
	title = fmt.Sprintf("%s (%s)", items["title"][config.GetLanguageCode()], config.GetAppVersionFull())
	systray.SetTitle("")
	access_mx.Unlock()
	updateIcon()
}

func FillMenu() {
	access_mx.Lock()
	defer access_mx.Unlock()
	lang := config.GetLanguageCode()

	log.Info(logContext, "is rebuilding its menu items")

	// clear old handlers
	if isFilledOnce {
		systray.ResetMenu()
		itemsCleared <- true
	}

	// title entry
	mTitle = systray.AddMenuItem(title, "")
	mTitle.Disable()

	// instance connections
	systray.AddSeparator()
	for instanceId, inst := range config.GetInstances() {
		connected, exists := instanceIdMapConnected[instanceId]
		if exists && connected {
			systray.AddMenuItem(fmt.Sprintf("%s:%d %s", inst.HostName, inst.HostPort, items["conYes"][lang]), "")
		} else {
			systray.AddMenuItem(fmt.Sprintf("%s:%d %s", inst.HostName, inst.HostPort, items["conNo"][lang]), "")
		}
	}

	// update available
	mUpdate = systray.AddMenuItem("-", "")
	mUpdate.Hide()
	if isUpdateAvailable {
		systray.AddSeparator()
		mUpdate.SetTitle(items["update"][lang])
		mUpdate.Show()
	}

	// last accessed files
	systray.AddSeparator()
	mFile0 = systray.AddMenuItem("-", "")
	mFile0.Hide()
	mFile1 = systray.AddMenuItem("-", "")
	mFile1.Hide()
	mFile2 = systray.AddMenuItem("-", "")
	mFile2.Hide()
	mFile3 = systray.AddMenuItem("-", "")
	mFile3.Hide()
	mFile4 = systray.AddMenuItem("-", "")
	mFile4.Hide()

	for i, f := range filesShow {
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

	// file open actions
	systray.AddSeparator()
	mConfig = systray.AddMenuItem(items["config"][lang], "")
	mLogs = systray.AddMenuItem(items["logs"][lang], "")

	// uninstall
	systray.AddSeparator()
	mUninstall0 = systray.AddMenuItem(items["uninstall0"][lang], "")
	mUninstall1 = mUninstall0.AddSubMenuItem(items["uninstall1"][lang], "")

	// toggle actions
	systray.AddSeparator()
	mDarkIcon = systray.AddMenuItemCheckbox(items["darkIcon"][lang], "", config.GetDarkIcon())
	mStartup = systray.AddMenuItemCheckbox(items["startup"][lang], "", config.GetAutoStart())
	mSsl = systray.AddMenuItemCheckbox(items["ssl"][lang], "", config.GetSsl())
	mSslVerify = systray.AddMenuItemCheckbox(items["sslVerify"][lang], "", config.GetSslVerify())
	if !config.GetSsl() {
		mSslVerify.Hide()
	}
	mDebug = systray.AddMenuItemCheckbox(items["debug"][lang], "", config.GetDebug())

	// quit
	systray.AddSeparator()
	mQuit = systray.AddMenuItem(items["quit"][lang], "")

	// handle menu item events
	isFilledOnce = true
	go func() {
		for {
			select {
			case <-itemsCleared:
				return
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
				open.WithLocalSystem(filepath.Join(config.GetPathApp(),
					config.GetFileName()), false)
			case <-mLogs.ClickedCh:
				open.WithLocalSystem(filepath.Join(config.GetPathApp(),
					config.GetFileNameLog()), false)
			case <-mDarkIcon.ClickedCh:
				config.SetDarkIcon(!config.GetDarkIcon())
				if err := config.WriteFile(); err != nil {
					continue
				}
				updateIcon()
				if config.GetDarkIcon() {
					mDarkIcon.Check()
				} else {
					mDarkIcon.Uncheck()
				}
			case <-mStartup.ClickedCh:
				config.SetAutoStart(!config.GetAutoStart())
				if err := config.WriteFile(); err != nil {
					continue
				}
				if err := install.AutoStart(); err != nil {
					continue
				}
				if config.GetAutoStart() {
					mStartup.Check()
				} else {
					mStartup.Uncheck()
				}
			case <-mSsl.ClickedCh:
				config.SetSsl(!config.GetSsl())
				if err := config.WriteFile(); err != nil {
					continue
				}
				if config.GetSsl() {
					mSsl.Check()
					mSslVerify.Show()
				} else {
					mSsl.Uncheck()
					mSslVerify.Hide()
				}
			case <-mSslVerify.ClickedCh:
				config.SetSslVerify(!config.GetSslVerify())
				if err := config.WriteFile(); err != nil {
					continue
				}
				if config.GetSslVerify() {
					mSslVerify.Check()
				} else {
					mSslVerify.Uncheck()
				}
			case <-mDebug.ClickedCh:
				config.SetDebug(!config.GetDebug())
				if err := config.WriteFile(); err != nil {
					continue
				}
				log.SetDebug(config.GetDebug())
				if config.GetDebug() {
					mDebug.Check()
				} else {
					mDebug.Uncheck()
				}
			case <-mUpdate.ClickedCh:
				if err := install.Update(isUpdateAvailableInstanceId); err != nil {
					log.Error(logContext, "failed to install update", err)
				}
			case <-mUninstall1.ClickedCh:
				if err := install.Remove(); err != nil {
					log.Error(logContext, "failed to uninstall", err)
				}
				config.OsExit <- syscall.SIGTERM
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

func updateIcon() {
	access_mx.Lock()
	defer access_mx.Unlock()

	darkIcon := config.GetDarkIcon()

	// 1st prio: any instance not connected
	for _, connected := range instanceIdMapConnected {
		if !connected {
			if darkIcon {
				systray.SetIcon(dark.Down)
			} else {
				systray.SetIcon(icon.Down)
			}
			return
		}
	}

	// 2nd prio: uploading
	if isLoadingUp.Load() {
		if darkIcon {
			systray.SetIcon(dark.Upload)
		} else {
			systray.SetIcon(icon.Upload)
		}
		return
	}

	// 3rd prio: downloading
	if isLoadingDown.Load() {
		if darkIcon {
			systray.SetIcon(dark.Download)
		} else {
			systray.SetIcon(icon.Download)
		}
		return
	}

	// 4th prio: action done
	if isActionDone.Load() {
		if darkIcon {
			systray.SetIcon(dark.Load)
		} else {
			systray.SetIcon(icon.Load)
		}
		return
	}

	// if nothing else: neutral
	if darkIcon {
		systray.SetIcon(dark.Neutral)
	} else {
		systray.SetIcon(icon.Neutral)
	}
}
