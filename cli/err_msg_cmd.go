package cli

import tea "github.com/charmbracelet/bubbletea"

type (
	errMsg error
)

func fwdError(err error) tea.Cmd {
	return func() tea.Msg {
		return errMsg(err)
	}
}
