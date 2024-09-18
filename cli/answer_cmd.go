package cli

import (
	"context"

	"github.com/anne-markis/fermtrack/answer"
	tea "github.com/charmbracelet/bubbletea"
)

type qAnswer string

type AiModel struct {
	ctx             context.Context
	Err             error
	answerClient    answer.AnsweringClient
	currentQuestion string
	currentAnswer   string
	// questionAnswered bool
	response string
}

func NewAIModel(client answer.AnsweringClient) AiModel {
	return AiModel{
		ctx:          context.Background(),
		answerClient: client,
	}
}

func (m AiModel) Update(msg tea.Msg) (AiModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			answer, err := m.answerClient.AskQuestion(m.ctx, m.currentQuestion)
			if err != nil {
				return m, fwdError(err)
			}
			m.currentAnswer = answer
			return m, giveAnswer(answer)
		}
	}
	return m, nil
}

func (m AiModel) View() string {
	return m.response
}

func (m *AiModel) Init() tea.Cmd {
	answer, err := m.answerClient.AskQuestion(m.ctx, "I'm new here, what can you help me with? 1 sentence")
	if err != nil {
		return fwdError(err)
	}
	m.currentAnswer = answer // TODO why is this needed both times, why not just use qAnswer
	return giveAnswer(answer)
}

func (m AiModel) Answer() string {
	return m.currentAnswer
}

func (m *AiModel) SetQuestion(q string) {
	m.currentQuestion = q
}

func giveAnswer(answer string) tea.Cmd {
	return func() tea.Msg {
		return qAnswer(answer)
	}
}
