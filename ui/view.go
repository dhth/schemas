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
	m.constraintsTitleStyle.Background(INACTIVE_PANE_TITLE_COLOR)

	switch m.activePane {
	case tablesList:
		m.tablesList.Styles.Title.Background(ACTIVE_PANE_TITLE_COLOR)
	case columnDetails:
		m.columnDetailsTitleStyle.Background(ACTIVE_PANE_TITLE_COLOR)
	case tableConstraints:
		m.constraintsTitleStyle.Background(ACTIVE_PANE_TITLE_COLOR)
	}

	switch m.fullScreenPane {
	case true:
		switch m.activeRHSPane {
		case columnDetails:
			content = m.columnDetailsStyle.Copy().Width(m.terminalWidth).PaddingLeft(3).Render(m.columnDetailsTitleStyle.Render("Columns") + "\n\n" + m.columns.View())
		case tableConstraints:
			content = m.constraintsStyle.Copy().Width(m.terminalWidth).PaddingLeft(3).Render(m.constraintsTitleStyle.Render("Constraints") + "\n\n" + m.constraints.View())
		}
	case false:
		var rhsView string
		tableListView := m.tableListStyle.Render(m.tablesList.View())
		switch m.activeRHSPane {
		case columnDetails:
			rhsView = m.columnDetailsStyle.Render(m.columnDetailsTitleStyle.Render("Columns") + "\n\n" + m.columns.View())
		case tableConstraints:
			rhsView = m.constraintsStyle.Render(m.constraintsTitleStyle.Render("Constraints") + "\n\n" + m.constraints.View())
		}
		content = lipgloss.JoinHorizontal(lipgloss.Top,
			tableListView,
			rhsView,
		)
	}
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#7c6f64"))

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("tab: switch focus, 1: columns, 2: constraints")
	}

	footerStr := fmt.Sprintf("%s%s",
		modeStyle.Render("schemas"),
		helpMsg,
	)
	footer = footerStyle.Render(footerStr)

	return lipgloss.JoinVertical(lipgloss.Left,
		content,
		statusBar,
		footer,
	)
}
