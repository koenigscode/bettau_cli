package view

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func DeckDetailsView(s *state.ApplicationState) {
	var chosen state.State

	huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[state.State]().
				Options(
					huh.NewOption("Learn new words", state.Learn),
					huh.NewOption("Review words", state.Review),
					huh.NewOption("Import from file", state.Import),
					huh.NewOption("Back", state.SelectLanguage),
				).
				Title("What would you like to do?").
				Description(fmt.Sprintf("You selected %s", s.CurrentLanguage)).
				Value(&chosen),
		),
	).Run()

	// if chosen == "Back" {
	// 	state.CurrentState = SelectLanguageState
	// 	return
	// }

	s.CurrentState = chosen

}
