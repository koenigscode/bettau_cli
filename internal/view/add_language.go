package view

import (
	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
	"github.com/koenigscode/bettau_cli/pkg/deck"
)

func AddLanguageView(s *state.ApplicationState) {
	var newDeckName string
	if s.CurrentState == state.AddLanguage {
		huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What language would you like to add?").
					Value(&newDeckName),
			),
		).Run()
	}
	s.Decks = append(s.Decks, deck.Deck{Name: newDeckName})
	s.CurrentState = state.SelectLanguage
}
