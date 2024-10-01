package main

import (
	"strings"
)

const helpText = `options:
	help		Get a list of commands
	list		List current winemaking projects
	view [UID] 	View specific project notes
	edit [UID]	Edit specific project notes
	clear		Clear screen and any selections
	Or simply ask the wine wizard anything you like!
`

const (
	AskWineWizard     = "ask"
	ListFermentations = "list"
	ClearList         = "clear"
	ViewFermentation  = "view"
	EditFermentation  = "edit"
)

type UserCommand struct {
	Command string
	Extra   string
}

func GetCommand(rawInput string) (UserCommand, error) {
	inputs := strings.Split(rawInput, " ")

	if len(inputs) == 0 {
		return UserCommand{Command: AskWineWizard}, nil
	}

	command := inputs[0]

	switch command {
	case "help":
		return UserCommand{Command: "help", Extra: helpText}, nil
	case ListFermentations:
		return UserCommand{Command: ListFermentations}, nil
	case ClearList:
		return UserCommand{Command: ClearList}, nil
	case ViewFermentation:
		return UserCommand{Command: ViewFermentation}, nil
	case "edit":
		return UserCommand{Command: EditFermentation}, nil
	}

	return UserCommand{Command: AskWineWizard}, nil
}
