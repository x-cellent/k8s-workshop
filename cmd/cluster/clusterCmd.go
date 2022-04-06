package cluster

import (
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/down"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/run"
)

var Cmd = &cobra.Command{
	Use:   "cluster",
	Short: "Runs the workshop cluster or exercises",
}

func init() {
	Cmd.AddCommand(run.Cmd)
	Cmd.AddCommand(down.Cmd)
}
