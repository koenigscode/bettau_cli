package view

import (
	"github.com/charmbracelet/huh"
	"github.com/koenigscode/bettau_cli/internal/state"
)

func ImportView(s *state.ApplicationState) {
	huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Title("Select a file:").
				AllowedTypes([]string{".lang"}).
				CurrentDirectory("/Users/koenig"),
		),
	).Run()
}
