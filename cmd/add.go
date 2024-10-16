package cmd

import (
	"errors"
	"fmt"

	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/model"
	"github.com/manifoldco/promptui"
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
	prompts := []promptui.Prompt{
		promptui.Prompt{
			Label:    "host",
			Validate: isNotBlank,
		},
		promptui.Prompt{
			Label:    "username",
			Validate: isNotBlank,
		},
		promptui.Prompt{
			HideEntered: true,
			Label:       "password",
		},
	}

	var results []string

	for _, prompt := range prompts {
		result, _ := prompt.Run()

		results = append(results, result)
	}

	host, user, pass := results[0], results[1], results[2]

	driverSelect := promptui.Select{
		Items: db.Drivers,
		Label: "driver",
	}

	_, driver, _ := driverSelect.Run()

	fmt.Println(driver)

	db.LocalDB.Create(&model.Database{
		Host: host,
		Pass: pass,
		User: user,
	})
}

func isNotBlank(val string) error {
	if len(val) > 0 {
		return nil
	}

	return errors.New("value cannot be blank")
}
