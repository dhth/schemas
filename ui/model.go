package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/schemas/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type activePane uint

const (
	tablesList activePane = iota
	columnDetails
)

type model struct {
	dbPool                  *pgxpool.Pool
	tablesList              list.Model
	columns                 table.Model
	columnsCache            map[string][]types.ColumnDetails
	message                 string
	messages                []string
	terminalHeight          int
	terminalWidth           int
	tableListStyle          lipgloss.Style
	columnDetailsStyle      lipgloss.Style
	columnDetailsTitleStyle lipgloss.Style
	activePane              activePane
}

func (m model) Init() tea.Cmd {
	return fetchTables(m.dbPool)
}
