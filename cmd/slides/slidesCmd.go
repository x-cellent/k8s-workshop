package slides

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/slides/flag"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const defaultPort = 8080

var Cmd = &cobra.Command{
	Use:   "slides",
	Short: fmt.Sprintf("Serves slides at http://localhost:%d/docs by default", defaultPort),
	RunE:  runServer,
}

func init() {
	Cmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}

	Cmd.PersistentFlags().IntP(flag.Port, flag.PortShort, defaultPort, "HTTP port of local workshop slides server")
}

func runServer(cmd *cobra.Command, args []string) error {
	docs := cmd.Context().Value("docs").(embed.FS)
	return Run("/docs", http.FS(docs))
}

func Run(path string, fs http.FileSystem) error {
	http.Handle("/", http.FileServer(fs))

	port := viper.GetInt(flag.Port)
	slidesURL := fmt.Sprintf("http://localhost:%d/docs", port)
	fmt.Printf("Starting workshop slides server at %s%s...\n", slidesURL, path)

	go func() {
		time.Sleep(500 * time.Millisecond)
		out, err := openInDefaultBrowser(slidesURL)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to open %s in default browser:\n%s%s", slidesURL, string(out), err.Error())
		}
	}()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func openInDefaultBrowser(fileOrURL string) ([]byte, error) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, fileOrURL)
	return exec.Command(cmd, args...).CombinedOutput()
}
