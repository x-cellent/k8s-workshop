package run

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Kind = string

const (
	Docker Kind = "docker"
	K8s    Kind = "k8s"
)

type manifestFragment struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

type manifest struct {
	fragment *manifestFragment
	content  []byte
}

const (
	delim = "\n========================================"

	nsPattern = `---
apiVersion: v1
kind: Namespace
metadata:
  name: ex%d
`
)

func Exercise(exercises embed.FS, n int, kind Kind) error {
	if kind == K8s {
		kubeconfig := os.Getenv("KUBECONFIG")
		if filepath.Base(kubeconfig) != "k8s-workshop.kubeconfig" {
			return fmt.Errorf("Please set KUBECONFIG environment variable pointing to path of workshop cluster kubeconfig")
		}
	}

	ns := fmt.Sprintf("ex%d", n)

	exercisesDir := filepath.Join("exercises", kind)
	dd, err := exercises.ReadDir(exercisesDir)
	if err != nil {
		return err
	}
	exDir := ""
	for _, d := range dd {
		if d.IsDir() && strings.HasPrefix(d.Name(), ns) {
			exDir = filepath.Join(exercisesDir, d.Name())
			break
		}
	}
	if exDir == "" {
		return fmt.Errorf("Example %d not found", n)
	}

	ss, err := getExerciseScripts(exercises, exDir)
	if err != nil {
		return err
	}

	folder := ""
	if len(ss) > 0 {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		folder = filepath.Join(".", fmt.Sprintf("%s-ex%d", kind, n))
		err = os.RemoveAll(folder)
		if err != nil {
			return err
		}

		err = os.Mkdir(folder, 0700)
		if err != nil {
			return err
		}

		err = os.Chdir(folder)
		if err != nil {
			return err
		}
		fmt.Printf("Executing script(s) in local folder %s...\n", folder)

		script := "script.sh"
		defer os.Remove(script)

		for _, bb := range ss {
			err = os.WriteFile(script, bb, 0700)
			if err != nil {
				return err
			}

			err = exec.Command("./" + script).Run()
			if err != nil {
				return err
			}
		}
		_ = os.Remove(script)

		err = os.Chdir(wd)
		if err != nil {
			return err
		}

		ff, err := os.ReadDir(folder)
		if err != nil {
			return err
		}

		if len(ff) == 0 {
			_ = os.Remove(folder)
			folder = ""
		}
	}

	err = printExercise(exercises, exDir, n)
	if err != nil {
		return err
	}

	if folder != "" {
		fmt.Printf("---> Wechsel dafür ins Unterverzeichnis %q und benutze die dort hinterlegte(n) Datei(n)!\n", folder)
	}

	if kind == K8s {
		kubectl, err := exec.LookPath("kubectl")
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

		fmt.Printf("Deploying exercise manifests into namespace %s...\n", ns)

		manifestFile := "manifest.yaml"
		defer os.Remove(manifestFile)

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

			args := []string{"apply", "-f", manifestFile}
			if m.fragment.Kind != "Namespace" {
				args = append([]string{"-n", ns}, args...)
			}
			err = exec.Command(kubectl, args...).Run()
			if err != nil {
				return err
			}
		}
	}

	return runSolutionTimer(exercises, exDir, kind, n)
}

func printExercise(fs embed.FS, exampleDir string, n int) error {
	bb, err := fs.ReadFile(filepath.Join(exampleDir, "exercise.md"))
	if err != nil {
		return err
	}

	ex := string(bb)
	if !strings.HasSuffix(ex, "\n\n") {
		ex += "\n"
	}
	title := "Vorbereitung"
	if n > 0 {
		title = fmt.Sprintf("Aufgabe %d", n)
	}
	fmt.Printf("%s\n%s:%s\n\n%s\n", delim, title, delim, ex)

	return nil
}

func runSolutionTimer(fs embed.FS, exampleDir string, kind Kind, n int) error {
	bb, err := fs.ReadFile(filepath.Join(exampleDir, "solution.md"))
	if err != nil {
		return nil
	}

	lines := strings.Split(string(bb), "\n")
	if len(lines) == 0 {
		return nil
	}

	folder := ""
	bb, err = fs.ReadFile(filepath.Join(exampleDir, "solution.sh"))
	if err == nil {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		folder = filepath.Join(".", fmt.Sprintf("%s-sol%d", kind, n))
		err = os.RemoveAll(folder)
		if err != nil {
			return err
		}

		err = os.Mkdir(folder, 0700)
		if err != nil {
			return err
		}

		err = os.Chdir(folder)
		if err != nil {
			return err
		}

		script := "script.sh"
		defer os.Remove(script)
		err = os.WriteFile(script, bb, 0700)
		if err != nil {
			return err
		}

		err = exec.Command("./" + script).Run()
		if err != nil {
			return err
		}
		_ = os.Remove(script)

		err = os.Chdir(wd)
		if err != nil {
			return err
		}

		ff, err := os.ReadDir(folder)
		if err != nil {
			return err
		}

		if len(ff) == 0 {
			_ = os.Remove(folder)
			folder = ""
		}
	}

	d, err := time.ParseDuration(strings.TrimSpace(lines[0]))
	if err == nil {
		fmt.Printf("\nDie Lösung wird in %s angezeigt\n", d)
		time.Sleep(d)
		lines[0] = fmt.Sprintf("%s\nLösung:%s\n", delim, delim)
	} else {
		fmt.Println(err.Error())
		fmt.Println(fmt.Sprintf("%s\nLösung:%s", delim, delim))
		fmt.Println()
	}

	fmt.Println(strings.Join(lines, "\n"))

	if folder != "" {
		fmt.Printf("---> Die Lösungsdateien liegen im Unterverzeichnis %q\n", folder)
	}

	return nil
}

func getExerciseScripts(fs embed.FS, embeddedDir string) ([][]byte, error) {
	dd, err := fs.ReadDir(embeddedDir)
	if err != nil {
		return nil, err
	}

	var ss [][]byte

	for _, d := range dd {
		path := filepath.Join(embeddedDir, d.Name())

		if d.IsDir() {
			scripts, err := getExerciseScripts(fs, path)
			if err != nil {
				return nil, err
			}
			ss = append(ss, scripts...)
			continue
		}

		if filepath.Ext(path) != ".sh" || filepath.Base(path) == "solution.sh" {
			continue
		}

		bb, err := fs.ReadFile(path)
		if err != nil {
			return nil, err
		}

		ss = append(ss, bb)
	}

	return ss, nil
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
			manifests, err := getManifests(fs, path)
			if err != nil {
				return nil, err
			}
			mm = append(mm, manifests...)
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
