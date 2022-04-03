package slides

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "slides",
	Short: "Serves slides at localhost:8080/docs",
	RunE:  runServer,
}

func runServer(cmd *cobra.Command, args []string) error {
	docs := cmd.Context().Value("docs").(embed.FS)
	return Run("/docs", http.FS(docs))
}

func Run(path string, fs http.FileSystem) error {
	http.Handle("/", http.FileServer(fs))

	fmt.Printf("Starting workshop slides server at localhost:8080%s...\n", path)
	return http.ListenAndServe(":8080", nil)
}
