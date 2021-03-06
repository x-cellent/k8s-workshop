package main

import (
	"embed"
	"fmt"
	"github.com/x-cellent/k8s-workshop/cmd"
	"os"
)

var (
	//go:embed docs/*
	docs embed.FS

	//go:embed exercises/*
	exercises embed.FS
)

func main() {
	err := cmd.Execute(docs, exercises)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
