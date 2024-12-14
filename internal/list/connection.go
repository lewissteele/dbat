package list

import (
	"github.com/charmbracelet/huh"
	"github.com/lewissteele/dbat/internal/db"
)

func RenderConnectionSelection() string {
	var c string
	options := []huh.Option[string]{}

	for _, v := range db.UserDBNames() {
		options = append(options, huh.NewOption[string](v, v))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("database").Value(&c).Options(
				options...,
			),
		),
	)

	err := form.Run()

	if err != nil {
		panic(err)
	}

	return c
}
