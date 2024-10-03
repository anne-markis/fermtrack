package main

import (
	"fmt"
	"strings"
)

const helpText = `options:
	help		Get a list of commands
	list		List current winemaking projects
	view [UID] 	View specific project notes
	edit [UID]	Edit specific project notes
	clear		Clear screen and any selections
	login		Login to the system
	Or simply ask the wine wizard anything you like!
`

const (
	AskWineWizard     = "ask"
	ListFermentations = "list"
	ClearList         = "clear"
	ViewFermentation  = "view"
	EditFermentation  = "edit"
	Login             = "login"
)

type UserCommand struct {
	Command string
	Extra   string
}

func GetCommand(rawInput string) (UserCommand, error) {
	inputs := strings.Split(rawInput, " ") // Note: this is pretty brittle. Stronger tokenization would be better

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
	case EditFermentation:
		return UserCommand{Command: EditFermentation}, nil
	case Login:
		cmd := UserCommand{Command: Login}
		if len(inputs) != 3 {
			return cmd, fmt.Errorf("usage: login username password")
		}
		return cmd, nil
	}

	return UserCommand{Command: AskWineWizard}, nil
}
