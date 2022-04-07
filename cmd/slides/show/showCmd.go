package show

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-cellent/k8s-workshop/cmd/open"
	"github.com/x-cellent/k8s-workshop/cmd/slides/show/flag"
	"net/http"
	"os"
	"time"
)

const defaultPort = 8080

var Cmd = &cobra.Command{
	Use:   "show",
	Short: fmt.Sprintf("Serves slides at http://localhost:%d/docs by default", defaultPort),
	RunE:  runServer,
}

func init() {
	Cmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(Cmd.PersistentFlags())
	}

	Cmd.PersistentFlags().IntP(flag.Port, flag.PortShort, defaultPort, "HTTP port of local workshop slides server")
	Cmd.PersistentFlags().StringP(flag.Jump, flag.JumpShort, "", "jump to slide: horizontal_slide/vertical_slide/fragment, e.g. '2' or '3/1' or '4/3/2'")
}

func runServer(cmd *cobra.Command, args []string) error {
	docs := cmd.Context().Value("docs").(embed.FS)
	return Run("/docs", viper.GetInt(flag.Port), http.FS(docs), true, viper.GetString(flag.Jump))
}

func Run(path string, port int, fs http.FileSystem, openInBrowser bool, jump string) error {
	http.Handle("/", http.FileServer(fs))

	if openInBrowser {
		if jump != "" {
			jump = "/#/" + jump
		}
		slidesURL := fmt.Sprintf("http://localhost:%d%s%s", port, path, jump)
		fmt.Printf("Starting workshop slides server at %s...\n", slidesURL)

		go func() {
			time.Sleep(500 * time.Millisecond)
			out, err := open.InDefaultBrowser(slidesURL)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Failed to open %s in default browser:\n%s%s", slidesURL, string(out), err.Error())
			}
		}()
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
