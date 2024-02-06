package main

import (
	"errors"
	"log"
	"os"

	"github.com/koenigscode/bettau_cli/internal/state"
	"github.com/koenigscode/bettau_cli/internal/view"
	"github.com/koenigscode/bettau_cli/pkg/deck"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	s := state.ApplicationState{CurrentState: state.SelectLanguage}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "learn",
				Aliases: []string{"l"},
				Usage:   "Learn a deck",
				Action: func(cCtx *cli.Context) error {
					filePath := cCtx.Args().First()
					if filePath == "" {
						// fmt.Println("Please provide a file path")
						return errors.New("Please provide a file path")
					}

					dat, err := os.ReadFile(filePath)
					if err != nil {
						panic(err)
					}

					d := deck.Deck{}
					yaml.Unmarshal(dat, &d)
					s.CurrentState = state.Learn
					s.CurrentDeck = d
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
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
