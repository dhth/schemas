package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/schemas/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitialModel(dbPool *pgxpool.Pool) model {

	stackItems := make([]list.Item, 0)

	var appDelegateKeys = newAppDelegateKeyMap()
	appDelegate := newAppItemDelegate(appDelegateKeys)

	// column details
	colTblCols := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Data Type", Width: 30},
		{Title: "Nullable", Width: 8},
	}
	columnsTable := table.New(
		table.WithColumns(colTblCols),
		table.WithFocused(true),
		table.WithHeight(8),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#83a598")).
		Bold(true)
	columnsTable.SetStyles(s)

	columnsCache := make(map[string][]types.ColumnDetails)

	// constraints
	constraintsTblCols := []table.Column{
		{Title: "Name", Width: 30},
		{Title: "Type", Width: 12},
		{Title: "Check Clause", Width: 40},
	}
	constraintsTbl := table.New(
		table.WithColumns(constraintsTblCols),
		table.WithFocused(true),
		table.WithHeight(8),
	)

	constraintsTbl.SetStyles(s)

	constraintsCache := make(map[string][]types.TableConstraint)

	baseStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color("#282828"))

	tableListStyle := baseStyle.
		PaddingTop(1).
		PaddingRight(2).
		PaddingLeft(1).
		PaddingBottom(1)

	columnDetailsStyle := lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		PaddingTop(1)

	columnDetailsTitleStyle := baseStyle.
		Bold(true)

	m := model{
		dbPool:                  dbPool,
		tablesList:              list.New(stackItems, appDelegate, 0, 0),
		columns:                 columnsTable,
		columnsCache:            columnsCache,
		constraints:             constraintsTbl,
		constraintsCache:        constraintsCache,
		tableListStyle:          tableListStyle,
		columnDetailsStyle:      columnDetailsStyle,
		columnDetailsTitleStyle: columnDetailsTitleStyle,
		constraintsStyle:        columnDetailsStyle,
		constraintsTitleStyle:   columnDetailsTitleStyle,
		activeRHSPane:           columnDetails,
		showHelp:                true,
	}
	m.tablesList.Title = "Tables"
	m.tablesList.SetStatusBarItemName("table", "tables")
	m.tablesList.DisableQuitKeybindings()
	m.tablesList.SetShowHelp(false)
	m.tablesList.Styles.Title = m.tablesList.Styles.Title.Foreground(lipgloss.Color("#282828")).Bold(true)

	return m
}
