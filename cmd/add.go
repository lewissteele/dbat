package cmd

import (
	"github.com/charmbracelet/huh"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Run:   run,
		Short: "save database connection",
		Use:   "add",
	})
}

func run(cmd *cobra.Command, args []string) {
	var (
		driver string
		host   string
		name   string
		pass   string
		port   string
		user   string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("driver").Value(&driver).Options(
				huh.NewOption(string(db.MariaDB), string(db.MariaDB)),
				huh.NewOption(string(db.MySQL), string(db.MySQL)),
				huh.NewOption(string(db.PostgreSQL), string(db.PostgreSQL)),
				huh.NewOption(string(db.SQLite), string(db.SQLite)),
			),
			huh.NewInput().Title("host").Value(&host).Validate(isNotBlank),
			huh.NewInput().Title("user").Value(&user).Validate(isNotBlank),
			huh.NewInput().Title("pass").Value(&pass),
			huh.NewInput().Title("port").Value(&port).Validate(isNotBlank),
			huh.NewInput().Title("name").Value(&name).Validate(isNotBlank),
		),
	)

	err := form.Run()

	if err != nil {
		panic(err)
	}

	db.LocalDB.Create(&model.Database{
		Driver: driver,
		Host:   host,
		Name:   name,
		Pass:   pass,
		Port:   port,
		User:   user,
	})
}
