package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color("#282828"))

	modeStyle = baseStyle.Copy().
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#b8bb26"))
)
