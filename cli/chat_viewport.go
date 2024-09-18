package cli

import "github.com/charmbracelet/bubbles/viewport"

func chatViewport() viewport.Model {
	viewPort := viewport.New(80, 5)
	viewPort.SetContent(`🍷🧙 Ask the wine wizard anything you like.`)
	return viewPort
}
