package slides

import (
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/slides/export"
	"github.com/x-cellent/k8s-workshop/cmd/slides/show"
)

var Cmd = &cobra.Command{
	Use:   "slides",
	Short: "Shows or exports workshop slides",
}

func init() {
	Cmd.AddCommand(show.Cmd)
	Cmd.AddCommand(export.Cmd)
}
