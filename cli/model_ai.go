package cli

import (
	"context"

	"github.com/anne-markis/fermtrack/answer"
	tea "github.com/charmbracelet/bubbletea"
)

type AnswerModel struct {
	ctx             context.Context
	Err             error
	answerClient    answer.AnsweringClient
	currentQuestion string
	// questionAnswered bool
	response string
}

func NewAnswerModel(client answer.AnsweringClient) AnswerModel {
	return AnswerModel{
		ctx:          context.Background(),
		answerClient: client,
	}
}

// need to break up the KeyEnter
// KeyEnter -> return 'startSpinner'
// spinStarted -> AskAQuestion
// giveAnser -> hide spinner
func (m AnswerModel) Update(msg tea.Msg) (AnswerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, setThinking(true)
		}
	case qIsThinking:
		if msg {
			answer, err := m.answerClient.AskQuestion(m.ctx, m.currentQuestion)
			if err != nil {
				return m, fwdError(err)
			}
			return m, tea.Batch(giveAnswer(answer), setThinking(false))
		}
	}
	return m, nil
}

func (m AnswerModel) View() string {
	return m.response
}

func (m *AnswerModel) Init() tea.Cmd {
	answer, err := m.answerClient.AskQuestion(m.ctx, "I'm new here, what can you help me with? 1 sentence")
	if err != nil {
		return fwdError(err)
	}
	return giveAnswer(answer)
}

func (m *AnswerModel) SetQuestion(q string) {
	m.currentQuestion = q
}
