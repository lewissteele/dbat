package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lewissteele/dbat/internal/db"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to saved connection",
	Run: func(cmd *cobra.Command, args []string) {
		host := getHost(args)
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

func getHost(args []string) string {
	if len(args) > 1 {
		return args[0]
	}

	prompt := promptui.Select{
		Label: "database",
		Items: []string{"localhost"},
	}
	
	_, host, err := prompt.Run()

	if err != nil {
		fmt.Println(err.Error())
	}

	return host
}
