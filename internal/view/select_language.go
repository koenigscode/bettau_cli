package view

import (
	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func SelectDeckView(s *state.ApplicationState) {
	const addLanguage = "Add language"
	const exit = "Exit"

	var chosen string
	languagesStr := make([]string, len(s.Decks))
	for i, d := range s.Decks {
		languagesStr[i] = d.Name
	}

	huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a language").
				Options(
					huh.NewOptions[string](append(languagesStr, addLanguage, exit)...)...,
				).
				Value(&chosen).
				WithKeyMap(huh.NewDefaultKeyMap()),
		),
	).Run()

	switch chosen {
	case addLanguage:
		s.CurrentState = state.AddLanguage
		return
	case exit:
		s.CurrentState = state.Exit
		return
	}

	for _, d := range s.Decks {
		if d.Name == chosen {
			s.CurrentState = state.DeckDetails
			s.CurrentDeck = d
			return
		}
	}

}
