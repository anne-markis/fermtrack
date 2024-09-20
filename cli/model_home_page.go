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
	isThinking       bool
	answerClient     answer.AnsweringClient
	senderStyle      lipgloss.Style
	responderStyle   lipgloss.Style
	err              error
}

func NewHomePage(aiAnswerer answer.AnsweringClient) homePage {
	return homePage{
		questionTextArea: questionTextArea(),
		responseViewPort: chatViewport(),
		thinkingSpinner:  thinkingSpinner(),
		answerClient:     aiAnswerer,
		senderStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		responderStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		err:              nil, // TODO use this

	}
}

func (h homePage) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, h.thinkingSpinner.Tick)
}

func (h homePage) View() string {
	var spinner string
	if h.isThinking {
		spinner = h.thinkingSpinner.View()
	}
	title := FermTrack_ANSIShadow()
	view := title + "\n" +
		lipgloss.JoinHorizontal(lipgloss.Top, h.questionTextArea.View(), spinner, h.responseViewPort.View()) +
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
			cmds = append(cmds, askQuestion(h.answerClient, h.questionTextArea.Value()), setThinking(true))

			h.questionTextArea.Reset()
			h.responseViewPort.GotoBottom()
		}
	case someAnswer:
		h.responseViewPort.SetContent(h.responderStyle.Render(fmt.Sprintf("🍷🧙: %v", msg)))
		h.questionTextArea.Reset()
		h.responseViewPort.GotoBottom()
		cmds = append(cmds, setThinking(false))
	case qIsThinking:
		if msg {
			h.isThinking = true
			h.responseViewPort.SetContent(h.responderStyle.Render(""))
		} else {
			h.isThinking = false
		}
	case errMsg:
		h.err = msg
		return h, nil
	default:
		var spinnerCmd tea.Cmd
		h.thinkingSpinner, spinnerCmd = h.thinkingSpinner.Update(msg)
		cmds = append(cmds, spinnerCmd)
	}

	return h, tea.Batch(cmds...)
}

func helpView() string {
	return helpStyle("\n\nctrl+c: Quit\n")
}

func chatViewport() viewport.Model {
	viewPort := viewport.New(80, 6)
	viewPort.SetContent(`🍷🧙 Ask the wine wizard anything you like.`)
	return viewPort
}

func questionTextArea() textarea.Model {
	textArea := textarea.New()
	textArea.Placeholder = "Ask away..."
	textArea.Focus()

	textArea.Prompt = "┃ "
	textArea.CharLimit = 300

	textArea.SetWidth(40)
	textArea.SetHeight(6)

	textArea.FocusedStyle.CursorLine = lipgloss.NewStyle()

	textArea.ShowLineNumbers = false
	textArea.KeyMap.InsertNewline.SetEnabled(false)
	return textArea
}

func thinkingSpinner() spinner.Model {
	thinkingSpinner := spinner.New()
	thinkingSpinner.Style = spinnerStyle
	thinkingSpinner.Spinner = spinner.Points
	return thinkingSpinner
}
