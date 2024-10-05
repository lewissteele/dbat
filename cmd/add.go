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
	prompt := promptui.Prompt{
		Label:    "host",
		Validate: isNotBlank,
	}

	host, _ := prompt.Run()

	dump.P(host)
}

func isNotBlank(val string) error {
	if len(val) > 0 {
		return nil
	}

	return errors.New("value cannot be blank")
}
