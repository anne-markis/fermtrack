package cli

import (
	"fmt"

	"github.com/anne-markis/fermtrack/answer"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type homePage struct {
	responseViewPort viewport.Model
	questionTextArea textarea.Model
	thinkingSpinner  spinner.Model
	aiAnswerer       AnswerModel
	senderStyle      lipgloss.Style
	responderStyle   lipgloss.Style
	err              error
}

func NewHomePage(aiAnswerer answer.AnsweringClient) homePage {
	return homePage{
		questionTextArea: questionTextArea(),
		responseViewPort: chatViewport(),
		thinkingSpinner:  thinkingSpinner(),
		aiAnswerer:       NewAnswerModel(aiAnswerer),
		senderStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		responderStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		err:              nil,
	}
}

func (h homePage) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, h.thinkingSpinner.Tick)
}

func (h homePage) View() string {
	title := FermTrack_ANSIShadow()
	view := title + "\n" + //h.thinkingSpinner.View() +
		lipgloss.JoinHorizontal(lipgloss.Top, h.questionTextArea.View(), h.responseViewPort.View()) +
		helpView()
	return view
}

func (h homePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd      tea.Cmd
		vpCmd      tea.Cmd
		aiCmd      tea.Cmd
		spinnerCmd tea.Cmd
	)

	cmds := []tea.Cmd{}

	h.questionTextArea, tiCmd = h.questionTextArea.Update(msg)
	h.responseViewPort, vpCmd = h.responseViewPort.Update(msg)
	h.thinkingSpinner, spinnerCmd = h.thinkingSpinner.Update(msg)

	cmds = append(cmds, tiCmd, vpCmd, aiCmd, spinnerCmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return h, tea.Quit
		case tea.KeyEnter:
			userText := h.questionTextArea.Value()
			h.aiAnswerer.SetQuestion(userText)

			h.questionTextArea.Reset()
			h.responseViewPort.GotoBottom()
		}
	case qAnswer:
		h.responseViewPort.SetContent(h.responderStyle.Render(fmt.Sprintf("🍷🧙: %v", msg)))
		h.questionTextArea.Reset()
		h.responseViewPort.GotoBottom()
	case qIsThinking:
		if msg {
			h.responseViewPort.SetContent(h.responderStyle.Render(fmt.Sprintf("🍷🧙: %v", h.thinkingSpinner.View())))
		}
	case errMsg:
		h.err = msg
		return h, nil
	default:

	}

	h.aiAnswerer, aiCmd = h.aiAnswerer.Update(msg) // this is what is holidng the UI up
	cmds = append(cmds, aiCmd)

	return h, tea.Batch(cmds...)
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

func thinkingSpinner() spinner.Model {
	thinkingSpinner := spinner.New()
	thinkingSpinner.Style = spinnerStyle
	thinkingSpinner.Spinner = spinner.Meter
	return thinkingSpinner
}
