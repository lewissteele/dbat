package cmd

import (
	"fmt"
	"os"

	"image/color"

	"github.com/c-bata/go-prompt"
	"github.com/charmbracelet/huh"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/lewissteele/dbat/internal/input"
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

func executor(query string) {
	if len(query) == 0 {
		return
	}

	if query == "exit" {
		os.Exit(0)
	}

	rows, err := db.Conn.Raw(query).Rows()

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

	err = db.Conn.ScanRows(rows, &results)

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

	go db.SaveHistory(query)
}

func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func selectedDB(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	var name string

	options := []huh.Option[string]{}

	for _, n := range db.UserDBNames() {
		options = append(options, huh.NewOption[string](n, n))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("database").Value(&name).Options(
				options...
			),
		),
	)

	err := form.Run()

	if err != nil {
		panic(err)
	}

	return name
}
