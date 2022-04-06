package k8s

import (
	"embed"
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/flag"
	"github.com/x-cellent/k8s-workshop/cmd/exercise/run"
)

var Cmd = &cobra.Command{
	Use:   "k8s",
	Short: "Runs the given Kubernetes exercise",
	RunE:  runExercise,
}

func runExercise(cmd *cobra.Command, args []string) error {
	n, err := cmd.Flags().GetInt(flag.Number)
	if err != nil {
		return err
	}
	fs := cmd.Context().Value("exercises").(embed.FS)
	return run.Exercise(fs, n, run.K8s)
}
