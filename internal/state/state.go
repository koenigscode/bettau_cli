package state

import "github.com/koenigscode/bettau_cli/pkg/deck"

type State int

const (
	Learn          State = iota
	DeckDetails    State = iota
	Review         State = iota
	Import         State = iota
	AddLanguage    State = iota
	SelectLanguage State = iota
	Exit           State = iota
)

type ApplicationState struct {
	CurrentState State
	CurrentDeck  deck.Deck
	Decks        []deck.Deck
}
