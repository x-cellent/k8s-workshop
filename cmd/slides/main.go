//go:build exclude

package main

import (
	"github.com/x-cellent/k8s-workshop/cmd/slides"
	"log"
)

func main() {
	log.Fatal(slides.RunStatic())
}
