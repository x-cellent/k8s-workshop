package exercise

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/docker"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/flag"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/k8s"
)

var Cmd = &cobra.Command{
	Use:              "exercise",
	Aliases:          []string{"ex"},
	Short:            "Runs the given exercise",
	TraverseChildren: true,
}

func init() {
	Cmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}

	Cmd.PersistentFlags().IntP(flag.Number, flag.NumberShort, 0, "exercise number")
	Cmd.PersistentFlags().Bool(flag.ShowSolution, false, "show solution immediately")

	_ = Cmd.PersistentFlags().MarkHidden(flag.ShowSolution)

	Cmd.AddCommand(docker.Cmd)
	Cmd.AddCommand(k8s.Cmd)
}
