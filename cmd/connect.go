package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to saved connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		db := db.GetUserDB(host)
		sql := prompt.Input("> ", completer)

		rows, err := db.Raw(sql).Rows()

		if err != nil {
			fmt.Println(err.Error())
		}

		columns, _ := rows.Columns()
		var headers table.Row

		for _, column := range columns {
			headers = append(headers, column)
		}

		t := table.NewWriter()
		t.AppendHeader(headers)

		var results []map[string]interface{}

		db.ScanRows(rows, &results)

		for idx, result := range results {
			if idx == 0 {
				continue
			}

			var values []interface{}

			for _, value := range result {
				values = append(values, value)
			}

			t.AppendRow(values)
		}

		fmt.Println(t.Render())
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}
