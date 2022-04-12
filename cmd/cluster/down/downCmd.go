package down

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/cluster"
	"os/exec"
)

var Cmd = &cobra.Command{
	Use:     "down",
	Aliases: []string{"shutdown", "stop", "rm", "remove", "del", "delete"},
	Short:   "Shutdown the workshop cluster",
	RunE:    shutdown,
}

func shutdown(cmd *cobra.Command, args []string) error {
	err := Shutdown()
	if err != nil {
		return err
	}

	fmt.Println("Workshop cluster deleted")

	return nil
}

func Shutdown() error {
	kind, err := exec.LookPath("kind")
	if err != nil {
		return err
	}

	err = exec.Command(kind, "delete", "cluster", "--name", cluster.Name).Run()
	if err != nil {
		return err
	}

	return nil
}
