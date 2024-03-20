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
			} else {
				return m, tea.Quit
			}
		case "tab", "shift+tab":
			if m.activePane == tablesList {
				m.activePane = columnDetails
				m.columnDetailsTitleStyle.Underline(true)
				m.tablesList.Styles.Title.Underline(false)
			} else {
				m.activePane = tablesList
				m.columnDetailsTitleStyle.Underline(false)
				m.tablesList.Styles.Title.Underline(true)
			}
		}

	case tea.WindowSizeMsg:
		w1, h1 := m.tableListStyle.GetFrameSize()
		w2, h2 := m.columnDetailsStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.tablesList.SetHeight(msg.Height - h1 - 2)
		m.tablesList.SetWidth(int(float64(msg.Width-w1) * 0.3))
		m.columns.SetHeight(msg.Height - h2 - 11)
		// 11 keeps the column table's lower end at the same level as the table list
		m.tablesList.SetWidth(int(float64(msg.Width-w2) * 0.7))
		m.tableListStyle = m.tableListStyle.Width(int(float64(msg.Width) * 0.3))
		m.columnDetailsStyle = m.columnDetailsStyle.Width(int(float64(msg.Width) * 0.7))
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
		cacheData, ok := m.columnsCache[msg.tableName]
		if !ok {
			return m, fetchTableDetails(m.dbPool, msg.tableName)
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
	case TablesDetailsFetchedMsg:
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
	}

	switch m.activePane {
	case tablesList:
		m.tablesList, cmd = m.tablesList.Update(msg)
		cmds = append(cmds, cmd)
	case columnDetails:
		m.columns, cmd = m.columns.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
