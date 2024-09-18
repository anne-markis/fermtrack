package cli

import (
	"fmt"

	"github.com/anne-markis/fermtrack/answer"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type homePage struct {
	viewport       viewport.Model
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
		aiAnswerer:     NewAIModel(aiAnswerer),
		senderStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		responderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		err:            nil,
	}
}

func (h homePage) Init() tea.Cmd {
	return textarea.Blink
}

func (h homePage) View() string {
	title := FermTrack_ANSIShadow()
	view := title + "\n" +
		lipgloss.JoinHorizontal(lipgloss.Top, h.textarea.View(), h.viewport.View()) +
		helpView()
	return view
}

func (h homePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		aiCmd tea.Cmd
	)

	h.textarea, tiCmd = h.textarea.Update(msg)
	h.viewport, vpCmd = h.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return h, tea.Quit
		case tea.KeyEnter:
			userText := h.textarea.Value()
			h.aiAnswerer.SetQuestion(userText)

			h.textarea.Reset()
			h.viewport.GotoBottom()
		}
	case qAnswer:
		h.viewport.SetContent(h.responderStyle.Render(fmt.Sprintf("🍷🧙: %v", msg)))
		h.textarea.Reset()
		h.viewport.GotoBottom()
	case errMsg:
		h.err = msg
		return h, nil
	}

	h.aiAnswerer, aiCmd = h.aiAnswerer.Update(msg)

	return h, tea.Batch(tiCmd, vpCmd, aiCmd)
}

func helpView() string {
	return helpStyle("\n\nctrl+c: Quit\n")
}

func chatViewport() viewport.Model {
	viewPort := viewport.New(80, 5)
	viewPort.SetContent(`🍷🧙 Ask the wine wizard anything you like.`)
	return viewPort
}

func questionTextArea() textarea.Model {
	textArea := textarea.New()
	textArea.Placeholder = "Ask away..."
	textArea.Focus()

	textArea.Prompt = "┃ "
	textArea.CharLimit = 280

	textArea.SetWidth(40)
	textArea.SetHeight(5)

	// Remove cursor line styling
	textArea.FocusedStyle.CursorLine = lipgloss.NewStyle()

	textArea.ShowLineNumbers = false
	textArea.KeyMap.InsertNewline.SetEnabled(false)
	return textArea
}
