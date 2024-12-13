package cmd

import (
	"fmt"
	"os"

	"image/color"

	"github.com/c-bata/go-prompt"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/input"
	"github.com/lewissteele/dbat/internal/output"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to saved database",
	Run: func(cmd *cobra.Command, args []string) {
		db.Connect(selectedDB(args))

		prompt := prompt.New(
			executor,
			input.Completer,
			prompt.OptionCompletionOnDown(),
			prompt.OptionHistory(db.History()),
			prompt.OptionSelectedSuggestionBGColor(prompt.Color(color.Black.Y)),
			prompt.OptionSelectedSuggestionTextColor(prompt.Color(color.White.Y)),
			prompt.OptionSuggestionBGColor(prompt.Black),
			prompt.OptionSuggestionTextColor(prompt.Color(color.White.Y)),
		)

		prompt.Run()
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func executor(q string) {
	if len(q) == 0 {
		return
	}

	if q == "exit" {
		os.Exit(0)
	}

	results, err := db.Query(q)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	output.RenderTable(results)
}

func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}
