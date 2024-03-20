package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var content string
	var footer string

	var statusBar string
	if m.message != "" {
		statusBar = RightPadTrim(m.message, m.terminalWidth)
	}

	tableListView := m.tableListStyle.Render(m.tablesList.View())
	tableDetailsView := m.columnDetailsStyle.Render(m.columnDetailsTitleStyle.Render("Columns") + "\n\n" + m.columns.View())

	content = lipgloss.JoinHorizontal(lipgloss.Top,
		tableListView,
		tableDetailsView,
	)
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
