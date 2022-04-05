package down

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/run"
	"os/exec"
)

var Cmd = &cobra.Command{
	Use:     "down",
	Aliases: []string{"shutdown", "stop", "rm", "remove", "del", "delete"},
	Short:   "Shutdown the workshop cluster",
	RunE:    shutdown,
}

func shutdown(cmd *cobra.Command, args []string) error {
	kind, err := exec.LookPath("kind")
	if err != nil {
		return err
	}

	err = exec.Command(kind, "delete", "cluster", "--name", run.ClusterName).Run()
	if err != nil {
		return err
	}

	docker, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	_ = exec.Command(docker, "network", "rm", run.NetworkName)

	fmt.Println("Workshop cluster deleted")

	return nil
}
