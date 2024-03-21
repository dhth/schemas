package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	INACTIVE_PANE_TITLE_COLOR = lipgloss.Color("#7c6f64")
	ACTIVE_PANE_TITLE_COLOR   = lipgloss.Color("#b8bb26")
)

func (m model) View() string {
	var content string
	var footer string

	var statusBar string
	if m.message != "" {
		statusBar = RightPadTrim(m.message, m.terminalWidth)
	}

	m.tablesList.Styles.Title.Background(INACTIVE_PANE_TITLE_COLOR)
	m.columnDetailsTitleStyle.Background(INACTIVE_PANE_TITLE_COLOR)

	switch m.activePane {
	case tablesList:
		m.tablesList.Styles.Title.Background(ACTIVE_PANE_TITLE_COLOR)
	case columnDetails:
		m.columnDetailsTitleStyle.Background(ACTIVE_PANE_TITLE_COLOR)
	}

	switch m.fullScreenPane {
	case true:
		content = m.columnDetailsStyle.Width(m.terminalWidth).PaddingLeft(3).Render(m.columnDetailsTitleStyle.Render("Columns") + "\n\n" + m.columns.View())
	case false:
		tableListView := m.tableListStyle.Render(m.tablesList.View())
		tableDetailsView := m.columnDetailsStyle.Render(m.columnDetailsTitleStyle.Render("Columns") + "\n\n" + m.columns.View())
		content = lipgloss.JoinHorizontal(lipgloss.Top,
			tableListView,
			tableDetailsView,
		)
	}
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#7c6f64"))

	footerStr := fmt.Sprintf("%s",
		modeStyle.Render("schemas"),
	)
	footer = footerStyle.Render(footerStr)

	return lipgloss.JoinVertical(lipgloss.Left,
		content,
		statusBar,
		footer,
	)
}
