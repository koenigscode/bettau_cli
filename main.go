package main

import (
	"flag"
	"os"

	"github.com/koenigscode/bettau_cli/internal/state"
	"github.com/koenigscode/bettau_cli/internal/view"
	"github.com/koenigscode/bettau_cli/pkg/deck"
	"gopkg.in/yaml.v3"
)

func main() {
	s := state.ApplicationState{CurrentState: state.SelectLanguage}

	var importFile = flag.String("i", "", "Path to deck to import")
	flag.Parse()

	if len(*importFile) > 0 {
		dat, err := os.ReadFile(*importFile)
		if err != nil {
			panic(err)
		}

		d := deck.Deck{}
		yaml.Unmarshal(dat, &d)
		s.CurrentState = state.DeckDetails
		s.CurrentDeck = d
	}

	// view.WelcomeView()

	for {
		switch s.CurrentState {
		case state.Learn:
			view.LearnView(&s)
		case state.SelectLanguage:
			view.SelectDeckView(&s)
		case state.AddLanguage:
			view.AddLanguageView(&s)
		case state.DeckDetails:
			view.DeckDetailsView(&s)
		case state.Import:
			view.ImportView(&s)
		case state.Exit:
			return
		}

	}

}
