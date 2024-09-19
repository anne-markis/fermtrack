package cli

import (
	tea "github.com/charmbracelet/bubbletea"
)

type qAnswer string

func giveAnswer(answer string) tea.Cmd {
	return func() tea.Msg {
		return qAnswer(answer)
	}
}
