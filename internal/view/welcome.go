package view

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func WelcomeView() {
	keymap := huh.KeyMap{
		Note: huh.NoteKeyMap{Submit: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "continue"))},
	}

	filePath := "./assets/ascii-art.txt"

	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Welcome to bettaU").
				Description(string(content)),
		),
	).WithKeyMap(&keymap).Run()
}
