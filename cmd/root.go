package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhth/schemas/internal/db"
	"github.com/dhth/schemas/internal/ui"
)

func Execute() error {
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

	dbPool, err := db.CreateDBPool()
	if err != nil {
		return err
	}

	defer dbPool.Close()

	return ui.RenderUI(dbPool)
}
