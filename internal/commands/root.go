package cmd

import (
	"github.com/dhth/schemas/internal/db"
	"github.com/dhth/schemas/internal/ui"
)

func Execute() error {
	dbPool, err := db.CreateDBPool()
	if err != nil {
		return err
	}

	defer dbPool.Close()

	return ui.RenderUI(dbPool)
}
