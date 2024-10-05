package cmd

import (
	"errors"

	"github.com/gookit/goutil/dump"
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

	dump.P(results)
}

func isNotBlank(val string) error {
	if len(val) > 0 {
		return nil
	}

	return errors.New("value cannot be blank")
}
