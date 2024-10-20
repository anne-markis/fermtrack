package main

import (
	"context"
	"fmt"
	"strings"

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

type userInfo struct {
	encodedToken string
	username     string
	uuid         string
}

type sessionState int

const (
	appView sessionState = iota
	entryVeiw
)

type mainModel struct {
	state sessionState

	responseViewPort   viewport.Model
	questionTextArea   textarea.Model
	thinkingSpinner    spinner.Model
	isThinking         bool
	fermTracker        client.Fermtracker
	fermentationsTable table.Model
	additionalHelp     string
	userInfo           userInfo
	ctx                context.Context
	err                error
}

func NewMainModel(fermTracker client.Fermtracker) mainModel {
	return mainModel{
		questionTextArea:   questionTextArea(),
		responseViewPort:   chatViewport(),
		thinkingSpinner:    thinkingSpinner(),
		fermTracker:        fermTracker,
		fermentationsTable: fermentationsTable([]client.Fermentation{}),
		ctx:                context.Background(),
		err:                nil, // TODO use this
	}
}

func StartCLI(ctx context.Context, fermTracker client.Fermtracker) error {
	p := tea.NewProgram(
		NewMainModel(fermTracker),
	)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (h mainModel) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, h.thinkingSpinner.Tick)
}

func (h mainModel) View() string {
	var spinner string
	if h.isThinking {
		spinner = h.thinkingSpinner.View()
	}

	title := FermTrack_ANSIShadow()

	defaultView := title + "\n" +
		lipgloss.JoinHorizontal(lipgloss.Top, h.questionTextArea.View(), spinner, h.responseViewPort.View(), h.fermentationsTable.View()) +
		h.helpView()

	// TODO nested models to vetter control flow/views
	switch h.state {
	case appView:
		return defaultView
	case entryVeiw:
		// TODO login screen
	}
	return defaultView
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd      tea.Cmd
		vpCmd      tea.Cmd
		aiCmd      tea.Cmd
		spinnerCmd tea.Cmd
		tableCmd   tea.Cmd
	)

	cmds := []tea.Cmd{}

	m.questionTextArea, tiCmd = m.questionTextArea.Update(msg)
	m.responseViewPort, vpCmd = m.responseViewPort.Update(msg)
	m.thinkingSpinner, spinnerCmd = m.thinkingSpinner.Update(msg)

	cmds = append(cmds, tiCmd, vpCmd, aiCmd, spinnerCmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			if m.fermentationsTable.Focused() {
				m.fermentationsTable.Blur()
				m.additionalHelp = "esc: select table"
			} else {
				m.fermentationsTable.Focus()
				m.additionalHelp = "esc: unselect table"
			}
		case tea.KeyEnter:
			input := m.questionTextArea.Value()

			systemCmd, err := GetCommand(input)
			if err != nil {
				cmds = append(cmds, fwdError(err))
			}

			switch systemCmd.Command {
			case AskWineWizard:
				cmds = append(cmds, m.askQuestion(input), setThinking(true))
			case ListFermentations:
				cmds = append(cmds, m.listFermentations(), setThinking(true))
			case ClearList:
				m.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙:")))
				m.fermentationsTable = fermentationsTable([]client.Fermentation{})
			case ViewFermentation:
				if m.fermentationsTable.SelectedRow() != nil {
					uuid := m.fermentationsTable.SelectedRow()[0] // uuid col is first col
					cmds = append(cmds, m.viewFermentation(uuid), setThinking(true))
				} else {
					m.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: Try using 'list' and selecting a row first!")))
				}
			case EditFermentation:
				if m.fermentationsTable.SelectedRow() != nil {
					// uuid := h.fermentationsTable.SelectedRow()[0] // uuid col is first col // TODO use https://github.com/charmbracelet/huh
					// cmds = append(cmds, h.startEditFermentation(uuid), setThinking(true))
					m.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: Unimplemented! %s", BubblyGlass())))
					m.responseViewPort.GotoTop()
				} else {
					m.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: Try using 'list' and selecting a row first!")))
				}
			case Login:
				var user string
				var pass string
				userInput := strings.Split(input, " ")
				if len(userInput) == 1 {
					// easy pass for meee
					user = "cooluser"
					pass = "cool123"
				} else {
					user = userInput[1]
					pass = userInput[2]
				}

				cmds = append(cmds, m.attemptLogin(user, pass), setThinking(true))
			default:
				m.responseViewPort.SetContent(responderStyle.Render(systemCmd.Extra)) // help text
			}

			m.questionTextArea.Reset()
		}
	case someAnswer:
		m.responseViewPort.SetContent(responderStyle.Render(fmt.Sprintf("🍷🧙: %v", msg)))
		m.questionTextArea.Reset()
		m.responseViewPort.GotoTop()
		cmds = append(cmds, setThinking(false))
	case qIsThinking:
		if msg {
			m.isThinking = true
			m.responseViewPort.SetContent(responderStyle.Render(""))
		} else {
			m.isThinking = false
		}
	case fermentationList:
		m.questionTextArea.Reset()

		m.fermentationsTable = fermentationsTable(msg.ferms)
		m.fermentationsTable.Focus()

		m.responseViewPort.SetContent(responderStyle.Render("🍷🧙: ok"))
		m.responseViewPort.GotoTop()
		cmds = append(cmds, setThinking(false))
	case fermentationView:
		m.responseViewPort.SetContent(responderStyle.Render(msg.ToString()))
		m.responseViewPort.GotoTop()
		cmds = append(cmds, setThinking(false))
	case fermentationStartEdit:
		// TODO
		cmds = append(cmds, setThinking(false))
	case userLogin:
		m.userInfo = userInfo{
			encodedToken: msg.Token,
			username:     msg.Username,
			uuid:         msg.UUID,
		}

		m.ctx = context.WithValue(m.ctx, client.ContextKeyJWT, m.userInfo.encodedToken)

		m.responseViewPort.SetContent(responderStyle.Render("🍷🧙: Login successful! Ask me anything or try `help`"))
		cmds = append(cmds, setThinking(false))
	case errMsg:
		m.err = msg
		return m, nil
	default:
		var spinnerCmd tea.Cmd
		m.thinkingSpinner, spinnerCmd = m.thinkingSpinner.Update(msg)
		cmds = append(cmds, spinnerCmd)
	}

	m.fermentationsTable, tableCmd = m.fermentationsTable.Update(msg)
	cmds = append(cmds, tableCmd)

	if m.userInfo.encodedToken == "" {
		m.questionTextArea.Placeholder = "login username password" // TODO: hit enter to bypass with default username/pass for ease of use
	} else {
		m.questionTextArea.Placeholder = "Ask away..."
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) helpView() string {
	helpText := "\n\n help: Get commands • ↑/↓: scroll answers • ctrl+c: Quit"
	if m.additionalHelp != "" {
		helpText = helpText + " • " + m.additionalHelp
	}

	return helpStyle(helpText)
}

func chatViewport() viewport.Model {
	viewPort := viewport.New(65, height)
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

func (m mainModel) askQuestion(q string) tea.Cmd {
	return func() tea.Msg {
		answer, err := m.fermTracker.AskQuestion(m.ctx, &client.FermentationQuestion{Question: q})
		if err != nil {
			return someAnswer(err.Error())
		}
		return someAnswer(answer.Answer)
	}
}

type fermentationList struct{ ferms []client.Fermentation }

func (m mainModel) listFermentations() tea.Cmd {
	return func() tea.Msg {
		ferms, err := m.fermTracker.ListFermentations(m.ctx)
		if err != nil {
			return someAnswer(err.Error())
		}
		return fermentationList{ferms}
	}
}

type fermentationView struct{ *client.Fermentation }

func (m mainModel) viewFermentation(uuid string) tea.Cmd {
	return func() tea.Msg {
		ferm, err := m.fermTracker.GetFermentation(m.ctx, uuid)
		if err != nil {
			return someAnswer(err.Error())
		}
		return fermentationView{ferm}
	}
}

type fermentationStartEdit struct{ *client.Fermentation }

func (m mainModel) startEditFermentation(uuid string) tea.Cmd { // TODO
	return func() tea.Msg {
		ferm, err := m.fermTracker.GetFermentation(m.ctx, uuid)
		if err != nil {
			return someAnswer(err.Error())
		}
		return fermentationStartEdit{ferm}
	}
}

type userLogin struct{ *client.LoginResponse }

func (m mainModel) attemptLogin(username, password string) tea.Cmd {
	return func() tea.Msg {
		resp, err := m.fermTracker.Login(m.ctx, username, password)
		if err != nil {
			return someAnswer(err.Error())
		}
		return userLogin{resp}
	}
}
