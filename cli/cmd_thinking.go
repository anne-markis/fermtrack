package cli

import tea "github.com/charmbracelet/bubbletea"

type qIsThinking bool

func setThinking(isThinking bool) tea.Msg {
	return func() tea.Msg {
		return qIsThinking(isThinking)
	}
}
