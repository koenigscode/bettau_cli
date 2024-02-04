package view

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func DeckDetailsView(s *state.ApplicationState) {
	var chosen state.State

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[state.State]().
				Title("What would you like to do?").
				Description(fmt.Sprintf("Current deck: %s", s.CurrentDeck.Name)).
				Options(
					huh.NewOption("Learn new words", state.Learn),
					huh.NewOption("Review words", state.Review),
					huh.NewOption("Import from file", state.Import),
					huh.NewOption("Back", state.SelectLanguage),
				).
				Value(&chosen),
		),
	)
	form.Run()

	s.CurrentState = chosen

}
