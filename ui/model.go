package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/schemas/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pane uint

const (
	tablesList Pane = iota
	columnDetails
	tableConstraints
	numPanes
)

type model struct {
	dbPool                  *pgxpool.Pool
	tablesList              list.Model
	columns                 table.Model
	columnsCache            map[string][]types.ColumnDetails
	constraints             table.Model
	constraintsCache        map[string][]types.TableConstraint
	message                 string
	messages                []string
	terminalHeight          int
	terminalWidth           int
	tableListStyle          lipgloss.Style
	columnDetailsStyle      lipgloss.Style
	columnDetailsTitleStyle lipgloss.Style
	constraintsStyle        lipgloss.Style
	constraintsTitleStyle   lipgloss.Style
	activePane              Pane
	lastPane                Pane
	activeRHSPane           Pane
	fullScreenPane          bool
	showHelp                bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		fetchTables(m.dbPool),
		hideHelp(time.Second*15),
	)
}
