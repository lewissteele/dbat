package cmd

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/lewissteele/dbat/internal/db"
)

func isNotBlank(val string) error {
	if len(val) > 0 {
		return nil
	}

	return errors.New("value cannot be blank")
}

func selectedDB(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	var name string

	options := []huh.Option[string]{}

	for _, n := range db.UserDBNames() {
		options = append(options, huh.NewOption[string](n, n))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("database").Value(&name).Options(
				options...,
			),
		),
	)

	err := form.Run()

	if err != nil {
		panic(err)
	}

	return name
}
