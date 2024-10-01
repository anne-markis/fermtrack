package main

import (
	"context"
	"fmt"

	"github.com/anne-markis/fermtrack/cli/client"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
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
	responseViewPort   viewport.Model
	questionTextArea   textarea.Model
	thinkingSpinner    spinner.Model
	isThinking         bool
	fermTracker        client.Fermtracker
	fermentationsTable table.Model
	additionalHelp     string
	err                error
}

func NewHomePage(fermTracker client.Fermtracker) homePage {
	return homePage{
		questionTextArea:   questionTextArea(),
		responseViewPort:   chatViewport(),
		thinkingSpinner:    thinkingSpinner(),
		fermTracker:        fermTracker,
		fermentationsTable: fermentationsTable([]client.Fermentation{}),
		err:                nil, // TODO use this
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
		lipgloss.JoinHorizontal(lipgloss.Top, h.questionTextArea.View(), spinner, h.responseViewPort.View(), h.fermentationsTable.View()) +
		h.helpView()
	return view
}

func (h homePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd      tea.Cmd
		vpCmd      tea.Cmd
		aiCmd      tea.Cmd
		spinnerCmd tea.Cmd
		tableCmd   tea.Cmd
	)

	cmds := []tea.Cmd{}

	h.questionTextArea, tiCmd = h.questionTextArea.Update(msg)
	h.responseViewPort, vpCmd = h.responseViewPort.Update(msg)
	h.thinkingSpinner, spinnerCmd = h.thinkingSpinner.Update(msg)

	cmds = append(cmds, tiCmd, vpCmd, aiCmd, spinnerCmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return h, tea.Quit
		case tea.KeyEsc:
			if h.fermentationsTable.Focused() {
				h.fermentationsTable.Blur()
				h.additionalHelp = "esc: select table"
			} else {
				h.fermentationsTable.Focus()
				h.additionalHelp = "esc: unselect table"
			}
		case tea.KeyEnter:
			input := h.questionTextArea.Value()

			systemCmd, err := GetCommand(input)
			if err != nil {
				cmds = append(cmds, fwdError(err))
			}

			switch systemCmd.Command {
			case AskWineWizard:
				cmds = append(cmds, askQuestion(h.fermTracker, input), setThinking(true))
			case ListFermentations:
				cmds = append(cmds, listFermentations(h.fermTracker), setThinking(true))
			case ClearList:
				h.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙:")))
				h.fermentationsTable = fermentationsTable([]client.Fermentation{})
			case ViewFermentation:
				if h.fermentationsTable.SelectedRow() != nil {
					uuid := h.fermentationsTable.SelectedRow()[0] // uuid col is first col
					cmds = append(cmds, viewFermentation(h.fermTracker, uuid), setThinking(true))
				} else {
					h.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: Try using 'list' and selecting a row first!")))
				}
			case EditFermentation:
				h.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: Unimplemented! %s", BubblyGlass())))
				h.responseViewPort.GotoTop()
			default:
				h.responseViewPort.SetContent(responderStyle.Render(systemCmd.Extra)) // help text
			}

			h.questionTextArea.Reset()
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
	case fermentationList:
		h.questionTextArea.Reset()

		h.fermentationsTable = fermentationsTable(msg.ferms)
		h.fermentationsTable.Focus()

		h.responseViewPort.SetContent(responderStyle.Render("🍷🧙: ok"))
		h.responseViewPort.GotoTop()
		cmds = append(cmds, setThinking(false))
	case fermentationView:
		h.responseViewPort.SetContent(responderStyle.Render(msg.ferm.ToString()))
		h.responseViewPort.GotoTop()
		cmds = append(cmds, setThinking(false))
	case errMsg:
		h.err = msg
		return h, nil
	default:
		var spinnerCmd tea.Cmd
		h.thinkingSpinner, spinnerCmd = h.thinkingSpinner.Update(msg)
		cmds = append(cmds, spinnerCmd)
	}

	h.fermentationsTable, tableCmd = h.fermentationsTable.Update(msg)
	cmds = append(cmds, tableCmd)

	return h, tea.Batch(cmds...)
}

func (h homePage) helpView() string {
	helpText := "\n\n help: Get commands • ↑/↓: scroll answers • ctrl+c: Quit"
	if h.additionalHelp != "" {
		helpText = helpText + " • " + h.additionalHelp
	}

	return helpStyle(helpText)
}

func chatViewport() viewport.Model {
	viewPort := viewport.New(60, height)
	viewPort.SetContent(`🍷🧙 Ask me, the wine wizard, anything you like.`)
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

func fermentationsTable(ferms []client.Fermentation) table.Model {
	columns := []table.Column{
		{Title: "UUID", Width: 36},
		{Title: "NickName", Width: 20},
		{Title: "Started At", Width: 15},
	}
	rows := make([]table.Row, len(ferms))
	for i, f := range ferms {
		rows[i] = table.Row{f.UUID, f.Nickname, f.StartAt.Format("2006-02-01")}
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
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
		answer, err := fermtrack.AskQuestion(context.Background(), &client.FermentationQuestion{Question: q})
		if err != nil {
			return someAnswer(err.Error())
		}
		return someAnswer(answer.Answer)
	}
}

type fermentationList struct{ ferms []client.Fermentation }

func listFermentations(fermtrack client.Fermtracker) tea.Cmd {
	return func() tea.Msg {
		ferms, err := fermtrack.ListFermentations(context.Background())
		if err != nil {
			return someAnswer(err.Error())
		}
		return fermentationList{ferms}
	}
}

type fermentationView struct{ ferm *client.Fermentation }

func viewFermentation(fermtrack client.Fermtracker, uuid string) tea.Cmd {
	return func() tea.Msg {
		ferm, err := fermtrack.GetFermentation(context.Background(), uuid)
		if err != nil {
			return someAnswer(err.Error())
		}
		return fermentationView{ferm}
	}
}
