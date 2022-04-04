package run

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var Cmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a kind cluster ready to be used for workshop exercises",
	RunE:  runCluster,
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

const (
	clusterName       = "k8s-workshop"
	clusterConfigFile = "k8s-workshop.kind.yaml"
	kubeconfigFile    = "k8s-workshop.kubeconfig"
)

func runCluster(cmd *cobra.Command, args []string) error {
	kind, err := exec.LookPath("kind")
	if err != nil {
		return err
	}

	docker, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	fmt.Println("Create Docker network named kind (10.10.10.0/24)")
	_ = exec.Command(docker, "network", "create", "kind", "--subnet", "10.10.10.0/24").Run()

	fmt.Printf("Write KinD config file %q\n", clusterConfigFile)
	err = os.WriteFile(clusterConfigFile, []byte(kindConfig), 0600)
	if err != nil {
		return err
	}

	fmt.Println("Starting cluster (this may take some time)...")
	err = exec.Command(kind, "create", "cluster", "--name", clusterName, "--config", clusterConfigFile).Run()
	if err != nil {
		return err
	}

	fmt.Printf("Write workshop cluster kubeconfig to %q\n", kubeconfigFile)
	bb, err := exec.Command(kind, "get", "kubeconfig", "--name", clusterName).CombinedOutput()
	if err != nil {
		return err
	}

	err = os.WriteFile(kubeconfigFile, bb, 0600)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	kubeconfigPath := filepath.Join(wd, kubeconfigFile)
	fmt.Printf("\nWorkshop cluster is ready to be used. Run\n\n    export KUBECONFIG=%s\n\nto autoconnect to the workshop cluster.\n", kubeconfigPath)

	return nil
}
