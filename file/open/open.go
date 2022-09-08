package open

import (
	"fmt"
	"os/exec"
	"runtime"
)

func WithLocalSystem(filePath string, choose bool) error {
	if choose {
		return chooseApp(filePath)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath)
	default:
		return fmt.Errorf("unsupported runtime environment '%v'", runtime.GOOS)
	}
	return cmd.Run()
}
