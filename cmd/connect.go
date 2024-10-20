package cmd

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/input"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"image/color"
)

var host string

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to saved database",
	Run: func(cmd *cobra.Command, args []string) {
		setHost(args)

		prompt := prompt.New(
			executor,
			input.Completer,
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

func executor(s string) {
	if len(s) == 0 {
		return
	}

	if s == "exit" {
		os.Exit(0)
	}

	userDB := db.UserDB(host)
	rows, err := userDB.Raw(s).Rows()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	columns, _ := rows.Columns()
	var headers table.Row

	for _, column := range columns {
		headers = append(headers, column)
	}

	t := table.NewWriter()
	t.AppendHeader(headers)

	var results []map[string]interface{}

	rows.Next()

	err = userDB.ScanRows(rows, &results)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, result := range results {
		var values []interface{}

		for _, c := range columns {
			values = append(values, result[c])
		}

		t.AppendRow(values)
	}

	fmt.Println(t.Render())

	go db.SaveHistory(s, host)
}

func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func setHost(args []string) {
	if len(args) > 0 {
		host = args[0]
		return
	}

	prompt := promptui.Select{
		HideHelp: true,
		Items:    db.UserDBNames(),
		Label:    "database",
	}

	_, host, _ = prompt.Run()
}
