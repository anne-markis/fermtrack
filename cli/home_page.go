package cli

import (
	"fmt"
	"strings"

	"github.com/anne-markis/fermtrack/answer"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type homePage struct {
	viewport       viewport.Model
	messages       []string // TODO only allow one question at a time, waiting for answer
	textarea       textarea.Model
	aiAnswerer     AiModel
	senderStyle    lipgloss.Style
	responderStyle lipgloss.Style
	err            error
}

func NewHomePage(aiAnswerer answer.AnsweringClient) homePage {
	return homePage{
		textarea:       questionTextArea(),
		viewport:       chatViewport(),
		messages:       []string{},
		aiAnswerer:     NewAIModel(aiAnswerer),
		senderStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		responderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		err:            nil,
	}
}

func (h homePage) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, h.aiAnswerer.Init())
}

// VIEW
// Entire UI display is a string
// So just like, keep that in mind
func (h homePage) View() string {
	title := FermTrack_ANSIShadow()
	view := title + "\n" +
		lipgloss.JoinHorizontal(lipgloss.Top, h.textarea.View(), h.viewport.View()) +
		helpView()
	return view
}

// UPDATE
// Handles user input
func (h homePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		aiCmd tea.Cmd
	)

	h.textarea, tiCmd = h.textarea.Update(msg)
	h.viewport, vpCmd = h.viewport.Update(msg)
	h.aiAnswerer, aiCmd = h.aiAnswerer.Update(msg) // might need to go after his handler Update

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return h, tea.Quit
		case tea.KeyEnter:
			userText := h.textarea.Value()
			h.messages = append(h.messages, h.senderStyle.Render("You: ")+userText)
			h.viewport.SetContent(strings.Join(h.messages, "\n"))
			h.aiAnswerer.SetQuestion(userText)
			h.textarea.Reset()
			h.viewport.GotoBottom()
		}
	case qAnswer:
		// h.messages = append(h.messages, h.responderStyle.Render("AI: ")+h.aiAnswerer.Answer())
		h.messages = append(h.messages, h.responderStyle.Render(fmt.Sprintf("AI: %v", msg)))
		h.viewport.SetContent(strings.Join(h.messages, "\n"))
		h.textarea.Reset()
		h.viewport.GotoBottom()
	case errMsg:
		h.err = msg
		return h, nil
	}

	// h.aiAnswerer, aiCmd = h.aiAnswerer.Update(msg) // might need to go after his handler Update

	return h, tea.Batch(tiCmd, vpCmd, aiCmd)
}

// https://charm.sh/blog/commands-in-bubbletea/

func fwdError(err error) tea.Cmd {
	return func() tea.Msg {
		return errMsg(err)
	}
}
