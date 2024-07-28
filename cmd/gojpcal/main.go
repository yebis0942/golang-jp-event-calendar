package main

import (
	"fmt"
	"os"

	gojpcal "github.com/yebis0942/golang-jp-event-calendar"
)

func main() {
	if err := gojpcal.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}
