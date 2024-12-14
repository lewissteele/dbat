package list

import (
	"github.com/charmbracelet/huh"
	"github.com/lewissteele/dbat/internal/db"
)

func RenderDatabaseSelection() string {
	var d string
	options := []huh.Option[string]{}

	for _, v := range db.Databases() {
		options = append(options, huh.NewOption[string](v, v))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("database").Value(&d).Options(
				options...,
			),
		),
	)

	err := form.Run()

	if err != nil {
		panic(err)
	}

	return d
}
