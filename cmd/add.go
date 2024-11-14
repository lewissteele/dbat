package cmd

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/gookit/goutil/dump"
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
		host string
		name string
		pass string
		port string
		user string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("driver").Value(&driver).Options(
				huh.NewOption(string(db.MariaDB), string(db.MariaDB)),
				huh.NewOption(string(db.MySQL), string(db.MySQL)),
				huh.NewOption(string(db.PostgreSQL), string(db.PostgreSQL)),
				huh.NewOption(string(db.SQLite), string(db.SQLite)),
			),
		),
	)

	err := form.Run()

	if err != nil {
		panic(err)
	}

	dump.P(driver)

	// selectDriver := promptui.Select{
	// 	Items: []db.Driver{
	// 		db.MariaDB,
	// 		db.MySQL,
	// 		db.PostgreSQL,
	// 		db.SQLite,
	// 	},
	// 	Label:    "driver",
	// 	HideHelp: true,
	// }
	//
	// _, driver, _ := selectDriver.Run()
	//
	// prompts := []promptui.Prompt{
	// 	promptui.Prompt{
	// 		Label:    "host",
	// 		Validate: isNotBlank,
	// 	},
	// 	promptui.Prompt{
	// 		Label:    "username",
	// 		Validate: isNotBlank,
	// 	},
	// 	promptui.Prompt{
	// 		Label: "password",
	// 		Mask:  '*',
	// 	},
	// 	promptui.Prompt{
	// 		AllowEdit: true,
	// 		Default:   db.Port(db.Driver(driver)),
	// 		Label:     "port",
	// 		Validate:  isNotBlank,
	// 	},
	// }

	// var results []string

	// for _, prompt := range prompts {
	// 	result, _ := prompt.Run()
	//
	// 	results = append(results, result)
	// }

	// host, user, pass, port := results[0], results[1], results[2], results[3]

	// promptName := promptui.Prompt{
	// 	AllowEdit: true,
	// 	Default:   host,
	// 	Label:     "name",
	// 	Validate:  isNotBlank,
	// }

	// name, _ := promptName.Run()
	//
	db.LocalDB.Create(&model.Database{
		Driver: driver,
		Host:   host,
		Name:   name,
		Pass:   pass,
		Port:   port,
		User:   user,
	})
}

func isNotBlank(val string) error {
	if len(val) > 0 {
		return nil
	}

	return errors.New("value cannot be blank")
}
