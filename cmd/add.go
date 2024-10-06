package cmd

import (
	"errors"

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

	host, username, password := results[0], results[1], results[2]

	db.LocalDB.Create(&model.Database{
		Host: host,
		Username: username,
		Password: password,
	})
}

func isNotBlank(val string) error {
	if len(val) > 0 {
		return nil
	}

	return errors.New("value cannot be blank")
}
