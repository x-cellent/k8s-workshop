//go:build exclude

package main

import (
	"github.com/x-cellent/k8s-workshop/cmd/slides/show"
	"log"
	"net/http"
)

func main() {
	log.Fatal(runStaticSlidesServer())
}

func runStaticSlidesServer() error {
	return show.Run("", 8080, http.Dir("docs"), true, "")
}
