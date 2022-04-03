package run

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the given cluster",
	RunE:  runCluster,
}

func runCluster(cmd *cobra.Command, args []string) error {
	return nil
}
