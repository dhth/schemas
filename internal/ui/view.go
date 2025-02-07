package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	inactivePaneTitleColor = lipgloss.Color("#7c6f64")
	activePaneTitleColor   = lipgloss.Color("#b8bb26")
)

func (m Model) View() string {
	var content string
	var footer string

	m.tablesList.Styles.Title = m.tablesList.Styles.Title.Background(inactivePaneTitleColor)
	m.columnDetailsTitleStyle = m.columnDetailsTitleStyle.Background(inactivePaneTitleColor)
	m.constraintsTitleStyle = m.constraintsTitleStyle.Background(inactivePaneTitleColor)

	switch m.activePane {
	case tablesList:
		m.tablesList.Styles.Title = m.tablesList.Styles.Title.Background(activePaneTitleColor)
	case columnDetails:
		m.columnDetailsTitleStyle = m.columnDetailsTitleStyle.Background(activePaneTitleColor)
	case tableConstraints:
		m.constraintsTitleStyle = m.constraintsTitleStyle.Background(activePaneTitleColor)
	}

	switch m.fullScreenPane {
	case true:
		switch m.activeRHSPane {
		case columnDetails:
			content = m.columnDetailsStyle.Width(m.terminalWidth).PaddingLeft(3).Render(m.columnDetailsTitleStyle.Render("Columns") + "\n\n" + m.columns.View())
		case tableConstraints:
			content = m.constraintsStyle.Width(m.terminalWidth).PaddingLeft(3).Render(m.constraintsTitleStyle.Render("Constraints") + "\n\n" + m.constraints.View())
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
		m.message,
		footer,
	)
}
