package kubectl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/download"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	"os"
	"runtime"
	"strings"
)

var Cmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Installs kubectl CLI into given destination directory",
	RunE:  install,
}

func install(cmd *cobra.Command, args []string) error {
	v := viper.GetString(flag.Version)

	if strings.ToLower(v) == "latest" {
		bb, err := download.Bytes("https://storage.googleapis.com/kubernetes-release/release/stable.txt")
		if err != nil {
			return err
		}
		v = string(bb)
	}

	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}

	fileURL := fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/%s/kubectl", v, runtime.GOOS, runtime.GOARCH)
	switch runtime.GOOS {
	case "windows":
		fileURL += ".exe"
	}

	path, err := download.File(fileURL, viper.GetString(flag.DestinationDir))
	if err != nil {
		return err
	}

	return os.Chmod(path, 0755)
}
