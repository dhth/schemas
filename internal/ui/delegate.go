package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type delegateKeyMap struct {
	choose key.Binding
}

func newAppDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("ctrl+f", "enter"),
			key.WithHelp("ctrl+f/enter", "check status"),
		),
	}
}

func newAppItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.
		SelectedTitle.
		Foreground(lipgloss.Color("#fe8019")).
		BorderLeftForeground(lipgloss.Color("#fe8019"))
	d.Styles.SelectedDesc = d.Styles.
		SelectedTitle

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		switch msgType := msg.(type) { // nolint:revive
		case tea.KeyMsg:
			switch { // nolint:revive
			case key.Matches(msgType,
				keys.choose,
				list.DefaultKeyMap().CursorUp,
				list.DefaultKeyMap().CursorDown,
				list.DefaultKeyMap().GoToStart,
				list.DefaultKeyMap().GoToEnd,
				list.DefaultKeyMap().NextPage,
				list.DefaultKeyMap().PrevPage):
				return chooseTableEntry(m.SelectedItem().FilterValue())
			}
		}

		return nil
	}

	return d
}
