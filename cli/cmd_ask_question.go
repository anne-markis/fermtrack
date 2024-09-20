package cli

import (
	"context"

	"github.com/anne-markis/fermtrack/answer"
	tea "github.com/charmbracelet/bubbletea"
)

type someAnswer string

func askQuestion(client answer.AnsweringClient, q string) tea.Cmd {
	return func() tea.Msg {
		answer, err := client.AskQuestion(context.Background(), q)
		if err != nil {
			return errMsg(err)
		}
		return someAnswer(answer)
	}
}
