package view

import (
	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func SelectLanguageView(s *state.ApplicationState) {
	const addLanguage = "Add language"
	const exit = "Exit"

	var chosen string
	languagesStr := make([]string, len(s.Languages))
	for i, lang := range s.Languages {
		languagesStr[i] = string(lang)
	}

	huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a language").
				Options(
					huh.NewOptions[string](append(languagesStr, addLanguage, exit)...)...,
				).
				Value(&chosen).
				WithKeyMap(huh.NewDefaultKeyMap()),
		),
	).Run()

	switch chosen {
	case addLanguage:
		s.CurrentState = state.AddLanguage
	case exit:
		s.CurrentState = state.Exit
	default:
		s.CurrentState = state.Learn
		s.CurrentLanguage = state.Language(chosen)
	}

}
