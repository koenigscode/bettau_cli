package view

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
	"github.com/koenigscode/bettau_cli/pkg/grader"
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
	).Run()

	g := grader.NewChatGPTGrader(os.Getenv("TOKEN"))
	gradeResult := g.Grade(grader.GradeQuery{Question: currentCard, Input: input})
	feedbackString := fmt.Sprintf("%s\n\nCorrect: %t\nSolution: %s\nFeedback: %s\n", input, gradeResult.Correct, gradeResult.Solution, gradeResult.Feedback)

	var chosenSelect string
	anotherQuestion := "Next question"
	exit := "Exit"
	huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title(currentCard).Description(feedbackString),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Options(
				huh.NewOptions[string](anotherQuestion, exit)...,
			).
				Value(&chosenSelect),
		),
	).Run()
	switch chosenSelect {
	case anotherQuestion:
		nextQuestion(s)
	case exit:
		s.CurrentState = state.DeckDetails
	}

}

func LearnView(s *state.ApplicationState) {
	nextQuestion(s)
}
