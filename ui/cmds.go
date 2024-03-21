package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dhth/schemas/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

func fetchTables(dbPool *pgxpool.Pool) tea.Cmd {
	return func() tea.Msg {
		tables, err := db.GetTables(dbPool)
		return TablesFetchedMsg{tables, err}
	}
}

func chooseTableEntry(tableName string) tea.Cmd {
	return func() tea.Msg {
		return TableChosenMsg{tableName}
	}
}

func fetchTableDetails(dbPool *pgxpool.Pool, tableName string) tea.Cmd {
	return func() tea.Msg {
		columns, err := db.GetColumnDetails(dbPool, tableName)
		return TableDetailsFetchedMsg{tableName, columns, err}
	}
}

func fetchTableConstraints(dbPool *pgxpool.Pool, tableName string) tea.Cmd {
	return func() tea.Msg {
		constraints, err := db.GetTableConstraints(dbPool, tableName)
		return TableConstraintsFetchedMsg{tableName, constraints, err}
	}
}

func hideHelp(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(time.Time) tea.Msg {
		return HideHelpMsg{}
	})
}
