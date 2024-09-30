package main

import (
	"context"
	"fmt"

	"github.com/anne-markis/fermtrack/cli/client"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	helpStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
	spinnerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	responderStyle = lipgloss.NewStyle()
	height         = 10
)

type homePage struct {
	responseViewPort viewport.Model
	questionTextArea textarea.Model
	thinkingSpinner  spinner.Model
	isThinking       bool
	fermTracker      client.Fermtracker
	err              error
}

func NewHomePage(fermTracker client.Fermtracker) homePage {
	return homePage{
		questionTextArea: questionTextArea(),
		responseViewPort: chatViewport(),
		thinkingSpinner:  thinkingSpinner(),
		fermTracker:      fermTracker,
		err:              nil, // TODO use this
	}
}

func StartCLI(ctx context.Context, fermTracker client.Fermtracker) error {
	p := tea.NewProgram(
		NewHomePage(fermTracker),
	)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
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
			input := h.questionTextArea.Value()

			systemCmd, _ := GetCommand(input) // TODO swallowed error
			if systemCmd == AskWineWizard {
				cmds = append(cmds, askQuestion(h.fermTracker, input), setThinking(true))
			} else {
				h.responseViewPort.SetContent(responderStyle.Render(systemCmd))
			}

			h.questionTextArea.Reset()
			h.responseViewPort.GotoBottom()
		}
	case someAnswer:
		h.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: %v", msg)))
		h.questionTextArea.Reset()
		h.responseViewPort.GotoTop()
		cmds = append(cmds, setThinking(false))
	case qIsThinking:
		if msg {
			h.isThinking = true
			h.responseViewPort.SetContent(responderStyle.Render(""))
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
	return helpStyle("\n\n help: Get commands • ↑/↓: scroll answers • ctrl+c: Quit\n")
}

func chatViewport() viewport.Model {
	viewPort := viewport.New(80, height)
	viewPort.SetContent(`🍷🧙 Ask me, the wine wizard, anything you like.`)
	// viewPort.Style = lipgloss.NewStyle()
	return viewPort
}

func questionTextArea() textarea.Model {
	textArea := textarea.New()
	textArea.Placeholder = "Ask away..."
	textArea.Focus()

	textArea.Prompt = "┃ "
	textArea.CharLimit = 300

	textArea.SetWidth(40)
	textArea.SetHeight(height)

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

type qIsThinking bool

func setThinking(isThinking bool) tea.Cmd {
	return func() tea.Msg {
		return qIsThinking(isThinking)
	}
}

type errMsg error

func fwdError(err error) tea.Cmd {
	return func() tea.Msg {
		return errMsg(err)
	}
}

type someAnswer string

func askQuestion(fermtrack client.Fermtracker, q string) tea.Cmd {
	return func() tea.Msg {
		answer, err := fermtrack.AskQuestion(context.Background(), q)
		if err != nil {
			// return errMsg(err)
			return someAnswer(err.Error())
		}
		return someAnswer(answer)
	}
}
