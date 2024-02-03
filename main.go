package main

import (
	"github.com/koenigscode/bettau_cli/internal/state"
	"github.com/koenigscode/bettau_cli/internal/view"
)

func main() {
	s := state.ApplicationState{CurrentState: state.SelectLanguage}

	view.WelcomeView()

	for {
		switch s.CurrentState {
		case state.SelectLanguage:
			view.SelectLanguageView(&s)
		case state.AddLanguage:
			view.AddLanguageView(&s)
		case state.Learn:
			view.DeckDetailsView(&s)
		case state.Import:
			view.ImportView(&s)
		case state.Exit:
			return
		}

	}

}
