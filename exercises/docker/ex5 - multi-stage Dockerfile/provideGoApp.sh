#!/usr/bin/env bash

cat <<EOF > app.go
package main

import "fmt"

func main() {
    fmt.Println("Hello Go!")
}
EOF
