package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.message = ""

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			fs := m.tablesList.FilterState()
			if fs == list.Filtering || fs == list.FilterApplied {
				m.tablesList.ResetFilter()
			} else if m.fullScreenPane {
				m.fullScreenPane = false
				m.activePane = m.lastPane
			} else {
				return m, tea.Quit
			}
		case "1":
			if !m.fullScreenPane {
				m.activeRHSPane = columnDetails
				if m.activePane != tablesList {
					m.activePane = columnDetails
				}
				selectedTable := m.tablesList.SelectedItem().FilterValue()
				_, ok := m.columnsCache[selectedTable]
				if !ok {
					cmds = append(cmds, fetchTableDetails(m.dbPool, selectedTable))
				}
			}
		case "2":
			if !m.fullScreenPane {
				m.activeRHSPane = tableConstraints
				if m.activePane != tablesList {
					m.activePane = tableConstraints
				}
				selectedTable := m.tablesList.SelectedItem().FilterValue()
				_, ok := m.constraintsCache[selectedTable]
				if !ok {
					cmds = append(cmds, fetchTableConstraints(m.dbPool, selectedTable))
				}
			}
		case "tab", "shift+tab":
			if m.activePane == tablesList {
				m.activePane = m.activeRHSPane
			} else {
				m.activePane = tablesList
			}
		case "ctrl+f":
			if !m.fullScreenPane {
				m.lastPane = m.activePane
				m.fullScreenPane = true
				switch m.activeRHSPane {
				case columnDetails:
					m.activePane = columnDetails
					m.columns.SetHeight(m.terminalHeight - 7)
				case tableConstraints:
					m.activePane = tableConstraints
					m.constraints.SetHeight(m.terminalHeight - 7)
				}
			} else {
				m.fullScreenPane = false
				m.activePane = m.lastPane
				switch m.activeRHSPane {
				case columnDetails:
					m.columns.SetHeight(m.terminalHeight - 12)
				case tableConstraints:
					m.constraints.SetHeight(m.terminalHeight - 12)
				}
			}
		}

	case HideHelpMsg:
		m.showHelp = false
	case tea.WindowSizeMsg:
		w1, h1 := m.tableListStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.tablesList.SetHeight(msg.Height - h1 - 2)
		m.tablesList.SetWidth(int(float64(msg.Width-w1) * 0.3))
		m.columns.SetHeight(msg.Height - 12)
		m.constraints.SetHeight(msg.Height - 12)
		// 12 keeps the column table's lower end at the same level as the table list
		m.tablesList.SetWidth(int(float64(msg.Width) * 0.7))
		m.tableListStyle = m.tableListStyle.Width(int(float64(msg.Width) * 0.3))
		m.columnDetailsStyle = m.columnDetailsStyle.Width(int(float64(msg.Width) * 0.7))
		m.constraintsStyle = m.constraintsStyle.Width(int(float64(msg.Width) * 0.7))
	case TablesFetchedMsg:
		if msg.err != nil {
			message := "error fetching tables: " + msg.err.Error()
			m.message = message
			m.messages = append(m.messages, message)
		} else {
			tableList := make([]list.Item, 0, len(msg.tables))
			for _, table := range msg.tables {
				tableList = append(tableList, table)
			}
			m.tablesList.SetItems(tableList)

			if len(msg.tables) > 0 {
				cmds = append(cmds, chooseTableEntry(msg.tables[0].Name))
			}
		}
	case TableChosenMsg:
		switch m.activeRHSPane {
		case columnDetails:
			cacheData, ok := m.columnsCache[msg.tableName]
			if !ok {
				cmds = append(cmds, fetchTableDetails(m.dbPool, msg.tableName))
			} else {
				var rows []table.Row
				for _, column := range cacheData {
					rows = append(rows,
						table.Row{
							column.Name,
							column.DataType,
							column.IsNullable,
						})
				}
				m.columns.SetRows(rows)
				m.columns.GotoTop()
			}
		case tableConstraints:
			cacheData, ok := m.constraintsCache[msg.tableName]
			if !ok {
				cmds = append(cmds, fetchTableConstraints(m.dbPool, msg.tableName))
			} else {
				var rows []table.Row
				for _, column := range cacheData {
					var checkClause string
					if column.CheckClause != nil {
						checkClause = *column.CheckClause
					}
					rows = append(rows,
						table.Row{
							column.Name,
							column.Type,
							checkClause,
						})
				}
				m.constraints.SetRows(rows)
				m.constraints.GotoTop()
			}
		}
	case TableDetailsFetchedMsg:
		if msg.err != nil {
			message := "error fetching table details: " + msg.err.Error()
			m.message = message
			m.messages = append(m.messages, message)
		} else {
			var rows []table.Row
			for _, column := range msg.columns {
				rows = append(rows,
					table.Row{
						column.Name,
						column.DataType,
						column.IsNullable,
					})
			}
			m.columns.SetRows(rows)
			m.columns.GotoTop()
			m.columnsCache[msg.tableName] = msg.columns
		}
	case TableConstraintsFetchedMsg:
		if msg.err != nil {
			message := "error fetching table details: " + msg.err.Error()
			m.message = message
			m.messages = append(m.messages, message)
		} else {
			var rows []table.Row
			for _, column := range msg.constraints {
				var checkClause string
				if column.CheckClause != nil {
					checkClause = *column.CheckClause
				}
				rows = append(rows,
					table.Row{
						column.Name,
						column.Type,
						checkClause,
					})
			}
			m.constraints.SetRows(rows)
			m.constraints.GotoTop()
			m.constraintsCache[msg.tableName] = msg.constraints
		}
	}

	switch m.activePane {
	case tablesList:
		m.tablesList, cmd = m.tablesList.Update(msg)
		cmds = append(cmds, cmd)
	case columnDetails:
		m.columns, cmd = m.columns.Update(msg)
		cmds = append(cmds, cmd)
	case tableConstraints:
		m.constraints, cmd = m.constraints.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
