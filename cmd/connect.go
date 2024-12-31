package cmd

import (
	"fmt"
	"os"

	"image/color"

	"github.com/c-bata/go-prompt"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/input"
	"github.com/lewissteele/dbat/internal/list"
	"github.com/lewissteele/dbat/internal/table"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to saved database",
	Run: func(cmd *cobra.Command, args []string) {
		var c string

		if len(args) > 0 {
			c = args[0]
		}

		if len(c) == 0 {
			c = list.RenderConnectionSelection()
		}

		db.Connect(c)

		if len(db.Selected()) == 0 && db.UserDB.Driver != string(db.SQLite) {
			db.Select(list.RenderDatabaseSelection())
		}

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

	if len(results) == 0 {
		return
	}

	table.Render(results)
	db.SaveHistory(q)
}
