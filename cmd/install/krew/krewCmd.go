package krew

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/download"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	kflag "github.com/x-cellent/k8s-workshop/cmd/install/krew/flag"
	"github.com/x-cellent/k8s-workshop/cmd/install/untar"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var Cmd = &cobra.Command{
	Use:   "krew",
	Short: "Installs krew plugin manager into given destination directory",
	RunE:  install,
}

func init() {
	Cmd.PersistentFlags().BoolP(kflag.Upgrade, kflag.UpgradeShort, false, "whether to upgrade kew if already installed")
}

func install(cmd *cobra.Command, args []string) error {
	v := viper.GetString(flag.Version)
	if strings.ToLower(v) == "latest" {
		v = "latest"
	} else if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}

	dir, err := os.MkdirTemp(".", "krew-*")
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

	krew := fmt.Sprintf("krew-%s_%s", runtime.GOOS, runtime.GOARCH)
	fileURL := fmt.Sprintf("https://github.com/kubernetes-sigs/krew/releases/%s/download/%s.tar.gz", v, krew)
	archive, err := download.File(fileURL, ".")
	if err != nil {
		return err
	}

	err = untar.Untar(archive, ".")
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "windows":
		krew = fmt.Sprintf(`\%s.exe`, krew)
	default:
		krew = "./" + krew
	}

	if viper.GetBool(kflag.Upgrade) {
		err = exec.Command(krew, "upgrade", "krew").Run()
	} else {
		err = exec.Command(krew, "install", "krew").Run()
	}
	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fmt.Printf("You need to add %s/.krew/bin to your PATH environment variable\n", home)

	return nil
}
