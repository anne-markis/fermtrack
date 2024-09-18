package cli

import "github.com/charmbracelet/lipgloss"

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

func helpView() string {
	return helpStyle("\n\nctrl+c: Quit\n") // TODO does q work
}
