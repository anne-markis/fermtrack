package main

import "strings"

const helpText = `options:
	help	Get a list of commands
	list	List current winemaking projects
	edit	Edit a project
	Or simply ask the wine wizard anything you like!
`

const AskWineWizard = "help" // TODO what?

func GetCommand(input string) (string, error) {
	input = strings.ReplaceAll(input, " ", "")

	switch input {
	case "help":
		return helpText, nil
	case "list":
		// TODO
	case "edit":
		// TODO
	}

	return AskWineWizard, nil
}
