package window

import (
	"strings"

	"github.com/go-vgo/robotgo"
)

func Focus(winTitle string) error {
	winTitle = strings.ToLower(winTitle)

	// go through all process window titles
	// FindIds() only goes through process names
	fpids, err := robotgo.Pids()
	if err != nil {
		return err
	}
	for _, fpid := range fpids {
		isExist, err := robotgo.PidExists(fpid)
		if err != nil {
			return err
		}
		if !isExist {
			continue
		}

		winTitlePid := strings.ToLower(robotgo.GetTitle(fpid))
		if winTitlePid != "" && strings.Contains(winTitlePid, winTitle) {
			return robotgo.ActivePid(fpid)
		}
	}

	// nothing found, attempt by window title
	return robotgo.ActiveName(winTitle)
}
