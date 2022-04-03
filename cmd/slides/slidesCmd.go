package slides

import (
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func RunStatic() error {
	return run("/", http.Dir("k8s-workshop"))
}

var Cmd = &cobra.Command{
	Use:   "slides",
	Short: "Serves slides at localhost:8080/k8s-workshop",
	RunE:  runServer,
}

var (
	//go:embed k8s-workshop/*
	slides embed.FS
)

func runServer(cmd *cobra.Command, args []string) error {
	return run("/k8s-workshop", http.FS(slides))
}

func run(path string, fs http.FileSystem) error {
	http.Handle("/", http.FileServer(fs))

	fmt.Printf("Starting workshop slides server at localhost:8080%s\n", path)
	return http.ListenAndServe(":8080", nil)
}
