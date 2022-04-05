package exercise

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/cluster/exercise/flag"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type manifestFragment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

type manifest struct {
	fragment *manifestFragment
	content  []byte
}

var Cmd = &cobra.Command{
	Use:     "exercise",
	Aliases: []string{"ex"},
	Short:   "Runs the given exercise",
	RunE:    runExercise,
}

const nsPattern = `---
apiVersion: v1
kind: Namespace
metadata:
  name: ex%d
`

func init() {
	Cmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}

	Cmd.PersistentFlags().IntP(flag.Number, flag.NumberShort, 0, "exercise number")

	_ = Cmd.MarkPersistentFlagRequired(flag.Number)
}

func runExercise(cmd *cobra.Command, args []string) error {
	kubectl, err := exec.LookPath("kubectl")
	if err != nil {
		return err
	}

	kubeconfig := os.Getenv("KUBECONFIG")
	if filepath.Base(kubeconfig) != "k8s-workshop.kubeconfig" {
		return fmt.Errorf("Please set KUBECONFIG environment variable pointing to path of workshop cluster kubeconfig")
	}

	exercises := cmd.Context().Value("exercises").(embed.FS)
	n := viper.GetInt(flag.Number)
	ns := fmt.Sprintf("ex%d", n)

	exDir := filepath.Join("exercises", ns)
	err = printExercise(exercises, exDir, n)
	if err != nil {
		return err
	}

	mm, err := getManifests(exercises, exDir)
	if err != nil {
		return err
	}

	if len(mm) == 0 {
		return nil
	}

	fmt.Println("Applying manifests...")

	manifestFile := "manifest.yaml"

	err = os.WriteFile(manifestFile, []byte(fmt.Sprintf(nsPattern, n)), 0600)
	if err != nil {
		return err
	}
	_ = exec.Command(kubectl, "delete", "-f", manifestFile, "--force", "--grace-period", "0").Run()
	err = exec.Command(kubectl, "apply", "-f", manifestFile).Run()
	if err != nil {
		return err
	}

	for _, m := range mm {
		err = os.WriteFile(manifestFile, m.content, 0600)
		if err != nil {
			return err
		}

		args = []string{"apply", "-f", manifestFile}
		if m.fragment.Kind != "Namespace" {
			args = append([]string{"-n", ns}, args...)
		}
		err = exec.Command(kubectl, args...).Run()
		if err != nil {
			return err
		}
	}

	return runSolutionTimer(exercises, exDir)
}

func printExercise(fs embed.FS, exampleDir string, number int) error {
	bb, err := fs.ReadFile(filepath.Join(exampleDir, "exercise.md"))
	if err != nil {
		return err
	}

	ex := string(bb)
	if !strings.HasSuffix(ex, "\n\n") {
		ex += "\n"
	}
	fmt.Printf("Aufgabe %d:\n\n%s\n", number, ex)

	return nil
}

func runSolutionTimer(fs embed.FS, exampleDir string) error {
	bb, err := fs.ReadFile(filepath.Join(exampleDir, "solution.md"))
	if err != nil {
		return err
	}

	lines := strings.Split(string(bb), "\n")
	if len(lines) == 0 {
		return nil
	}

	d, err := time.ParseDuration(strings.TrimSpace(lines[0]))
	if err == nil {
		fmt.Printf("\nDie Lösung wird in %s angezeigt\n", d)
		time.Sleep(d)
		lines[0] = "\nLösung:\n"
	} else {
		fmt.Println(err.Error())
		fmt.Println("\nLösung:")
		fmt.Println()
	}

	fmt.Println(strings.Join(lines, "\n"))

	return nil
}

func getManifests(fs embed.FS, embeddedDir string) ([]*manifest, error) {
	dd, err := fs.ReadDir(embeddedDir)
	if err != nil {
		return nil, err
	}

	var mm []*manifest

	for _, d := range dd {
		path := filepath.Join(embeddedDir, d.Name())

		if d.IsDir() {
			ff, err := getManifests(fs, path)
			if err != nil {
				return nil, err
			}
			mm = append(mm, ff...)
			continue
		}

		if filepath.Ext(path) != ".yaml" {
			continue
		}

		bb, err := fs.ReadFile(path)
		if err != nil {
			return nil, err
		}
		m := &manifest{
			content: bb,
		}
		err = yaml.Unmarshal(bb, &m.fragment)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		if m.fragment.Kind == "Namespace" {
			continue
		}

		mm = append(mm, m)
	}

	return mm, nil
}
