package main

import (
	"fmt"
	"os"

	"github.com/anne-markis/fermtrack/state"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 15
const defaultWidth = 20

var (
	titleStyle      = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

func main() {
	// TODO get crom categories
	items := []list.Item{
		state.Item("Wine"),
		state.Item("Kraut"),
		state.Item("Pickles"),
	}

	l := list.New(items, state.ItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Type of Ferm?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := state.AppState{List: l, ListType: "category"}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
