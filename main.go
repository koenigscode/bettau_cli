package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type Language string

const (
	LearnState          = "LearnState"
	ReviewState         = "ReviewState"
	ImportState         = "ImportState"
	AddLanguageState    = "AddLanguageState"
	SelectLanguageState = "SelectLanguageState"
	ExitState           = "ExitState"
)

const (
	BackAction        = "Back"
	ExitAction        = "Exit"
	AddLanguageAction = "Add language"
)

type State struct {
	CurrentState    string
	CurrentLanguage string
	Languages       []string
}

func SelectLanguage(state *State) {
	var chosen string
	huh.NewSelect[string]().
		Title("Choose a language").
		Options(
			huh.NewOptions[string](append(state.Languages, "Add language", "Exit")...)...,
		).
		Value(&chosen).
		WithKeyMap(huh.NewDefaultKeyMap()).
		Run()

	switch chosen {
	case AddLanguageAction:
		state.CurrentState = AddLanguageState
	case ExitAction:
		state.CurrentState = ExitState
	default:
		state.CurrentState = LearnState
		state.CurrentLanguage = chosen
	}

}
func AddLanguage(state *State) {
	var newLanguage string
	if state.CurrentState == AddLanguageState {
		huh.NewInput().
			Title("What language would you like to add?").
			Value(&newLanguage).
			Run()
	}
	state.Languages = append(state.Languages, newLanguage)
	state.CurrentState = SelectLanguageState
}

func LearnMenu(state *State) {
	var chosen string
	huh.NewSelect[string]().
		Options(
			huh.NewOption("Learn new words", LearnState),
			huh.NewOption("Review words", ReviewState),
			huh.NewOption("Import from file", ImportState),
			huh.NewOption("Back", BackAction),
		).
		Title("What would you like to do?").
		Description(fmt.Sprintf("You selected %s", state.CurrentLanguage)).
		Value(&chosen).
		Run()

	if chosen == "Back" {
		state.CurrentState = SelectLanguageState
		return
	}

	state.CurrentState = chosen

}
func ImportFromFile(state *State) {
	huh.NewFilePicker().
		Title("Select a file:").
		AllowedTypes([]string{".lang"}).
		CurrentDirectory("/Users/koenig").
		Run()
}

func main() {
	state := State{CurrentState: SelectLanguageState}

	huh.NewNote().
		Title("Welcome to lang-learn").
		Description("Helps you learn languages bla blabla").
		Run()

	for {
		switch state.CurrentState {
		case SelectLanguageState:
			SelectLanguage(&state)
		case AddLanguageState:
			AddLanguage(&state)
		case LearnState:
			LearnMenu(&state)
		case ImportState:
			ImportFromFile(&state)
		case ExitState:
			return
		}

	}

}
