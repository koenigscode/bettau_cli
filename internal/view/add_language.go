package view

import (
	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func AddLanguageView(s *state.ApplicationState) {
	var newLanguage string
	if s.CurrentState == state.AddLanguage {
		huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What language would you like to add?").
					Value(&newLanguage),
			),
		).Run()
	}
	s.Languages = append(s.Languages, state.Language(newLanguage))
	s.CurrentState = state.SelectLanguage
}
