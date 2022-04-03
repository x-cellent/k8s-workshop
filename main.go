package main

import (
	"fmt"
	"github.com/x-cellent/k8s-workshop/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
