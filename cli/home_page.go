package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

type homePage struct {
	// answerClient answer.AnsweringClient
	viewport       viewport.Model
	messages       []string
	textarea       textarea.Model
	senderStyle    lipgloss.Style
	responderStyle lipgloss.Style
	err            error
}

func NewHomePage() homePage {
	textArea := textarea.New()
	textArea.Placeholder = "Ask away..."
	textArea.Focus()

	textArea.Prompt = "┃ "
	textArea.CharLimit = 280

	textArea.SetWidth(30)
	textArea.SetHeight(5)

	// Remove cursor line styling
	textArea.FocusedStyle.CursorLine = lipgloss.NewStyle()

	textArea.ShowLineNumbers = false

	viewPort := viewport.New(50, 5)
	viewPort.SetContent(`🍷🧙 Ask the wine wizard anything you like.`)

	textArea.KeyMap.InsertNewline.SetEnabled(false)

	return homePage{
		textarea:       textArea,
		messages:       []string{},
		viewport:       viewPort,
		senderStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		responderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		err:            nil,
	}
}

func (h homePage) Init() tea.Cmd {
	return textarea.Blink
}

// VIEW
// Entire UI display is a string
// So just like, keep that in mind
func (h homePage) View() string {
	title := FermTrack_ANSIShadow()
	return fmt.Sprintf("%s\n%s\n\n%s%s", title, h.viewport.View(), h.textarea.View(), helpView())
}

// UPDATE
// Handles user input
func (h homePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	h.textarea, tiCmd = h.textarea.Update(msg)
	h.viewport, vpCmd = h.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(h.textarea.Value())
			return h, tea.Quit
		case tea.KeyEnter:
			h.messages = append(h.messages, h.senderStyle.Render("You: ")+h.textarea.Value())
			// m.messages = append(m.messages, m.responderStyle.Render("THEM: whatever"))
			h.viewport.SetContent(strings.Join(h.messages, "\n"))
			h.textarea.Reset()
			h.viewport.GotoBottom()
		}
	case errMsg:
		h.err = msg
		return h, nil
	}

	return h, tea.Batch(tiCmd, vpCmd)
}
