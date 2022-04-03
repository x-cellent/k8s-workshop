package helm

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/cp"
	"github.com/x-cellent/k8s-workshop/cmd/install/download"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	"github.com/x-cellent/k8s-workshop/cmd/install/untar"
	"github.com/x-cellent/k8s-workshop/cmd/install/unzip"
	"github.com/x-cellent/k8s-workshop/cmd/open"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "helm",
	Short: "Installs helm CLI into given destination directory",
	RunE:  install,
}

func install(cmd *cobra.Command, args []string) error {
	v := viper.GetString(flag.Version)

	if strings.ToLower(v) == "latest" {
		link := "https://github.com/helm/helm/tags"
		msg := fmt.Sprintf("Cannot find latest version of helm, please specify a version from %q", link)
		fmt.Println(msg)
		time.Sleep(3 * time.Second)
		_, _ = open.InDefaultBrowser(link)
		return fmt.Errorf("Version needed")
	}

	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}

	var suffix string
	switch runtime.GOOS {
	case "windows":
		suffix = "zip"
	default:
		suffix = "tar.gz"
	}

	dir, err := os.MkdirTemp(".", "helm-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Chdir(dir)
	if err != nil {
		return err
	}
	defer os.Chdir(currentDir)

	fileURL := fmt.Sprintf("https://get.helm.sh/helm-%s-%s-%s.%s", v, runtime.GOOS, runtime.GOARCH, suffix)
	archive, err := download.File(fileURL, viper.GetString(flag.DestinationDir))
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "windows":
		err = unzip.Unzip(archive, ".")
	default:
		err = untar.Untar(archive, ".")
	}
	if err != nil {
		return err
	}

	path := filepath.Join(currentDir, "helm")
	switch runtime.GOOS {
	case "windows":
		path += ".exe"
	}

	src := filepath.Join(fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH), path)
	dest := filepath.Join(viper.GetString(flag.DestinationDir), path)
	return cp.File(src, dest)
}
