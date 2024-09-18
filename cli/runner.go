package cli

import (
	"context"

	"github.com/anne-markis/fermtrack/answer"
	tea "github.com/charmbracelet/bubbletea"
)

func StartCLI(ctx context.Context, answerClient answer.AnsweringClient) error {
	p := tea.NewProgram(
		NewHomePage(answerClient),
	)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
