package docker

import (
	"embed"
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/flag"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/run"
)

var Cmd = &cobra.Command{
	Use:     "docker",
	Aliases: []string{"d"},
	Short:   "Runs the given Docker exercise",
	RunE:    runExercise,
}

func runExercise(cmd *cobra.Command, args []string) error {
	n, err := cmd.Flags().GetInt(flag.Number)
	if err != nil {
		return err
	}
	showSolution, err := cmd.Flags().GetBool(flag.ShowSolution)
	if err != nil {
		showSolution = false
	}
	fs := cmd.Context().Value("exercises").(embed.FS)
	return run.Exercise(fs, n, run.Docker, showSolution)
}
