package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RenderUI(dbPool *pgxpool.Pool) error {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			return err
		}
		defer f.Close()
	}

	p := tea.NewProgram(InitialModel(dbPool), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
