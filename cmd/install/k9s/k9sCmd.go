package k9s

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/download"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	"runtime"
	"strings"
	"syscall"
)

var Cmd = &cobra.Command{
	Use:   "k9s",
	Short: "Installs k9s CLI into given destination directory",
	RunE:  install,
}

func install(cmd *cobra.Command, args []string) error {
	v := viper.GetString(flag.Version)

	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}

	if strings.ToLower(v) == "latest" {
	}

	fileURL := fmt.Sprintf("", v, runtime.GOOS, runtime.GOARCH)
	switch runtime.GOOS {
	case "windows":
		fileURL += ".exe"
	}

	path, err := download.File(fileURL, viper.GetString(flag.DestinationDir))
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		return syscall.Chmod(path, 0755)
	default:
		return nil
	}
}
