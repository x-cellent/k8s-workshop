//go:build exclude

package main

import (
	"github.com/x-cellent/k8s-workshop/cmd/slides"
	"log"
	"net/http"
)

func main() {
	log.Fatal(slides.Run("", http.Dir("docs")))
}
