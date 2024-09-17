package state

import (
	"github.com/anne-markis/fermtrack/model"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO doesnt belong in state
var (
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type AppState struct {
	List           list.Model // category list
	CategoryChoice string     // current category

	FermTable table.Model

	Quitting bool
}

func (m AppState) Init() tea.Cmd {
	return nil
}

func (m AppState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.List.SelectedItem().(model.ListItem)
			if ok {
				m.CategoryChoice = string(i.Name)

			}
			// m.table, cmd = m.table.Update(msg)
			return m, tea.Quit // TODO return cmd
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m AppState) View() string {
	if m.CategoryChoice != "" {
		var filename string
		switch m.CategoryChoice {
		case "wine":
			filename = "./storage/wine.csv"
		case "kraut":
			filename = "./storage/kraut.csv"
		case "pickes":
			filename = "./storage/pickles.csv" // TODO doesn't load?
		}

		csvReader := CSVData{}
		if err := csvReader.ParseCSV(filename); err != nil {
			quitTextStyle.Render(err.Error())
			return ""
		}

		columns := []table.Column{}
		rows := []table.Row{}
		for _, col := range csvReader.Headers {
			columns = append(columns, table.Column{Title: col, Width: 20})
		}
		for _, row := range csvReader.Data {
			rows = append(rows, row)
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
		m.FermTable = t
		return m.FermTable.View()
	}
	if m.Quitting {
		return quitTextStyle.Render("Buh bye.")
	}
	return "\n" + m.List.View()
}
