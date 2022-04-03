package slides

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/open"
	"github.com/x-cellent/k8s-workshop/cmd/slides/flag"
	"net/http"
	"os"
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
	return Run("/docs", viper.GetInt(flag.Port), http.FS(docs))
}

func Run(path string, port int, fs http.FileSystem) error {
	http.Handle("/", http.FileServer(fs))

	slidesURL := fmt.Sprintf("http://localhost:%d%s", port, path)
	fmt.Printf("Starting workshop slides server at %s...\n", slidesURL)

	go func() {
		time.Sleep(500 * time.Millisecond)
		out, err := open.InDefaultBrowser(slidesURL)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to open %s in default browser:\n%s%s", slidesURL, string(out), err.Error())
		}
	}()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
