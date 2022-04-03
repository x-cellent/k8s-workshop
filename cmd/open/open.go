package open

import (
	"os/exec"
	"runtime"
)

func InDefaultBrowser(fileOrURL string) ([]byte, error) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, fileOrURL)
	return exec.Command(cmd, args...).CombinedOutput()
}
