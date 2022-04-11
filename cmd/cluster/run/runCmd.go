package run

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/cluster"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/down"
	"os"
	"os/exec"
	"path/filepath"
)

var Cmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"create", "up", "start"},
	Short:   "Runs a kind cluster ready to be used for workshop exercises",
	RunE:    runCluster,
}

const kindConfig = `---
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: ingress-ready=true
  extraPortMappings:
  - containerPort: 30000
    hostPort: 30000
    protocol: TCP
  - containerPort: 30001
    hostPort: 30001
    protocol: TCP
`

func runCluster(cmd *cobra.Command, args []string) error {
	err := down.Shutdown()
	if err != nil {
		return err
	}

	kind, err := exec.LookPath("kind")
	if err != nil {
		return err
	}

	docker, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	fmt.Printf("Create Docker network named %s (10.10.10.0/24)", cluster.NetworkName)
	_ = exec.Command(docker, "network", "create", cluster.NetworkName, "--subnet", "10.10.10.0/24").Run()

	fmt.Printf("Write KinD config file %q\n", cluster.ClusterConfigFile)
	err = os.WriteFile(cluster.ClusterConfigFile, []byte(kindConfig), 0600)
	if err != nil {
		return err
	}

	fmt.Println("Starting cluster (this may take some time)...")
	err = exec.Command(kind, "create", "cluster", "--name", cluster.ClusterName, "--config", cluster.ClusterConfigFile).Run()
	if err != nil {
		return err
	}

	fmt.Printf("Write workshop cluster kubeconfig to %q\n", cluster.KubeconfigFile)
	bb, err := exec.Command(kind, "get", "kubeconfig", "--name", cluster.ClusterName).CombinedOutput()
	if err != nil {
		return err
	}

	err = os.WriteFile(cluster.KubeconfigFile, bb, 0600)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	kubeconfigPath := filepath.Join(wd, cluster.KubeconfigFile)
	fmt.Printf("\nWorkshop cluster is ready to be used. Run\n\n    export KUBECONFIG=%s\n\nto autoconnect to the workshop cluster using kubectl, helm or k9s.\n", kubeconfigPath)

	return nil
}
