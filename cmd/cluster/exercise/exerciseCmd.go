package exercise

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/exercise/flag"
)

var Cmd = &cobra.Command{
	Use:     "exercise",
	Aliases: []string{"ex"},
	Short:   "Runs the given exercise",
	RunE:    runExercise,
}

func init() {
	Cmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}

	Cmd.PersistentFlags().IntP(flag.Number, flag.NumberShort, 0, "exercise number")

	_ = Cmd.MarkPersistentFlagRequired(flag.Number)
}

func runExercise(cmd *cobra.Command, args []string) error {
	e := viper.GetInt(flag.Number)
	_ = e
	return nil
}
