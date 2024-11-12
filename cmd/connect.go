package cmd

import (
	"fmt"
	"os"

	"image/color"

	"github.com/c-bata/go-prompt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/input"
	"github.com/lewissteele/dbat/internal/model"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var userDB *model.Database
var conn *gorm.DB

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to saved database",
	Run: func(cmd *cobra.Command, args []string) {
		userDB = db.UserDB(selectedDB(args))
		conn = userDB.Conn()

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

func executor(query string) {
	if len(query) == 0 {
		return
	}

	if query == "exit" {
		os.Exit(0)
	}

	rows, err := conn.Raw(query).Rows()

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

	err = conn.ScanRows(rows, &results)

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

	go db.SaveHistory(query, *userDB)
}

func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func selectedDB(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	prompt := promptui.Select{
		HideHelp: true,
		Items:    db.UserDBNames(),
		Label:    "database",
	}

	_, name, _ := prompt.Run()

	return name
}
