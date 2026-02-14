package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhth/schemas/internal/commands"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Inspect postgres schemas via a TUI.

schemas needs the following environment variables to be set:
- DATABASE_ADDRESS 
- DATABASE_PORT
- DATABASE_USERNAME
- DATABASE_PASSWORD
- DATABASE_DBNAME
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
