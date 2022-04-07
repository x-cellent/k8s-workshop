package export

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/slides/export/flag"
	"github.com/x-cellent/k8s-workshop/cmd/slides/show"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var Cmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"pdf"},
	Short:   "Exports workshop slides to PDF",
	RunE:    exportToPDF,
}

func init() {
	Cmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}

	Cmd.PersistentFlags().StringP(flag.DestinationPath, flag.DestinationPathShort, "./slides.pdf", "path to destination")
}

func exportToPDF(cmd *cobra.Command, args []string) error {
	docker, err := exec.LookPath("docker")
	if err != nil {
		return err
	}

	port := 8080
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		docs := cmd.Context().Value("docs").(embed.FS)
		wg.Done()
		err = show.Run("/docs", port, http.FS(docs), false, "")
	}()
	wg.Wait()
	time.Sleep(time.Second)
	if err != nil {
		return err
	}

	dest := viper.GetString(flag.DestinationPath)
	hostPath := filepath.Dir(dest)
	if strings.HasPrefix(hostPath, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		hostPath = home + hostPath[1:]
	} else if strings.HasPrefix(hostPath, ".") || !strings.HasPrefix(hostPath, string(filepath.Separator)) {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		if strings.HasPrefix(hostPath, ".") {
			hostPath = wd + hostPath[1:]
		} else {
			hostPath = filepath.Join(wd, hostPath)
		}
	}
	filename := filepath.Base(dest)
	vol := fmt.Sprintf("%s:/work", hostPath)
	url := fmt.Sprintf("http://localhost:%d/docs", port)
	out, err := exec.Command(docker, "run", "--rm", "-t", "--net", "host", "-v", vol, "astefanutti/decktape", url, filepath.Join("/work", filename)).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return exec.Command(docker, "rm", "-f", "k8s-slides").Run()
}
