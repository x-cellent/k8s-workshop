package install

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/install/flag"
	"github.com/x-cellent/k8s-workshop/cmd/install/helm"
	"github.com/x-cellent/k8s-workshop/cmd/install/k9s"
	"github.com/x-cellent/k8s-workshop/cmd/install/kind"
	"github.com/x-cellent/k8s-workshop/cmd/install/krew"
	"github.com/x-cellent/k8s-workshop/cmd/install/kubectl"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Installs tools on local machine",
}

func init() {
	defer func() {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}()

	Cmd.PersistentFlags().StringP(flag.Version, flag.VersionShort, "latest", "tool version")
	Cmd.PersistentFlags().StringP(flag.DestinationDir, flag.DestinationDirShort, ".", "destination directory")

	Cmd.AddCommand(kubectl.Cmd)
	Cmd.AddCommand(krew.Cmd)
	Cmd.AddCommand(helm.Cmd)
	Cmd.AddCommand(k9s.Cmd)
	Cmd.AddCommand(kind.Cmd)
}
