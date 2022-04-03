package k9s

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/cp"
	"github.com/x-cellent/k8s-workshop/cmd/install/download"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	"github.com/x-cellent/k8s-workshop/cmd/install/untar"
	"github.com/x-cellent/k8s-workshop/cmd/open"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "k9s",
	Short: "Installs k9s CLI into given destination directory",
	RunE:  install,
}

func install(cmd *cobra.Command, args []string) error {
	v := viper.GetString(flag.Version)

	if strings.ToLower(v) == "latest" {
		link := "https://github.com/derailed/k9s/tags"
		msg := fmt.Sprintf("Cannot find latest version of k9s, please specify a version from %q", link)
		fmt.Println(msg)
		time.Sleep(3 * time.Second)
		_, _ = open.InDefaultBrowser(link)
		return fmt.Errorf("Version needed")
	}

	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}

	dir, err := os.MkdirTemp(".", "k9s-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	destDir := viper.GetString(flag.DestinationDir)
	if destDir == "." {
		destDir = currentDir
	}

	err = os.Chdir(dir)
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)

	goos := strings.ToTitle(runtime.GOOS)
	arch := runtime.GOARCH
	switch arch {
	case "amd64":
		arch = "x86_64"
	}

	fileURL := fmt.Sprintf("https://github.com/derailed/k9s/releases/download/%s/k9s_%s_%s.tar.gz", v, goos, arch)
	archive, err := download.File(fileURL, ".")
	if err != nil {
		return err
	}

	err = untar.Untar(archive, ".")
	if err != nil {
		return err
	}

	path := "k9s"
	switch runtime.GOOS {
	case "windows":
		path += ".exe"
	}

	dest := filepath.Join(destDir, path)
	err = cp.File(path, dest)
	if err != nil {
		return err
	}

	return os.Chmod(dest, 0755)
}
