package ui

import (
	"charm.land/lipgloss/v2"
)

var (
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color("#282828"))

	modeStyle = baseStyle.
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#b8bb26"))

	helpMsgStyle = baseStyle.
			Bold(true).
			Foreground(lipgloss.Color("#83a598"))
)
