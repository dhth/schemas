package cmd

import (
	"github.com/dhth/schemas/db"
	"github.com/dhth/schemas/ui"
)

func Execute() {
	dbPool := db.CreateDBPool()
	defer dbPool.Close()

	ui.RenderUI(dbPool)
}
