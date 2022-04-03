package kind

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/cp"
	"github.com/x-cellent/k8s-workshop/cmd/install/download"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	"github.com/x-cellent/k8s-workshop/cmd/open"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "kind",
	Short: "Installs kind CLI into given destination directory",
	RunE:  install,
}

func install(cmd *cobra.Command, args []string) error {
	v := viper.GetString(flag.Version)

	if strings.ToLower(v) == "latest" {
		link := "https://github.com/kubernetes-sigs/kind/tags"
		msg := fmt.Sprintf("Cannot find latest version of kind, please specify a version from %q", link)
		fmt.Println(msg)
		time.Sleep(3 * time.Second)
		_, _ = open.InDefaultBrowser(link)
		return fmt.Errorf("Version needed")
	}

	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}

	fileURL := fmt.Sprintf("https://github.com/kubernetes-sigs/kind/releases/download/%s/kind-%s-%s", v, runtime.GOOS, runtime.GOARCH)
	switch runtime.GOOS {
	case "windows":
		fileURL += ".exe"
	}

	path, err := download.File(fileURL, viper.GetString(flag.DestinationDir))
	if err != nil {
		return err
	}
	defer os.Remove(path)

	dest := filepath.Join(viper.GetString(flag.DestinationDir), "kind")
	switch runtime.GOOS {
	case "windows":
		dest += ".exe"
	}

	err = cp.File(path, dest)
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		return os.Chmod(dest, 0755)
	default:
		return nil
	}
}
