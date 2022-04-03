//go:build exclude

package main

import (
	"github.com/x-cellent/k8s-workshop/cmd/slides"
	"log"
	"net/http"
)

func main() {
	log.Fatal(runStaticSlidesServer())
}

func runStaticSlidesServer() error {
	return slides.Run("", 8080, http.Dir("docs"))
}
