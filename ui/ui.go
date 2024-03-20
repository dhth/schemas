package ui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RenderUI(dbPool *pgxpool.Pool) {

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
	p := tea.NewProgram(InitialModel(dbPool), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Something went wrong %s", err)
	}
}
