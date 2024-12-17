package list

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/model"
)

func RenderConnectionSelection() string {
	var c string
	options := []huh.Option[string]{}

	var databases []model.Database
	db.LocalDB.Find(&databases)

	for _, d := range databases {
		key := fmt.Sprintf(
			"%s - %s@%s",
			d.Name,
			d.User,
			d.Host,
		)

		if d.Driver == string(db.SQLite) {
			key = fmt.Sprintf("%s - %s", d.Name, d.Path)
		}

		options = append(options, huh.NewOption[string](
			key,
			d.Name,
		))
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
