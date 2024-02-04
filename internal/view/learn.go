package view

import (
	"math/rand"

	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func nextQuestion(s *state.ApplicationState) {
	var currentCard string
	var input string

	currentCard = s.CurrentDeck.Contents[rand.Intn(len(s.CurrentDeck.Contents))]

	huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(currentCard).
				Value(&input),
		),
		huh.NewGroup(
			huh.NewConfirm().Title("Continue?"),
		).WithHideFunc(func() bool {
			return len(input) == 0
		}),
	).Run()

}

func LearnView(s *state.ApplicationState) {
	nextQuestion(s)
}
