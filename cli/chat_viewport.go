package cli

import "github.com/charmbracelet/bubbles/viewport"

func chatViewport() viewport.Model {
	// TODO make viewport dynamic? or put in different panel than the chat, it is being cut fof
	viewPort := viewport.New(80, 5) // second arg is the gap at the top
	viewPort.SetContent(`🍷🧙 Ask the wine wizard anything you like.`)
	return viewPort
}
