package table

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
)

func Render(d []map[string]interface{}) {
	var headers table.Row

	for k, _ := range d[0] {
		headers = append(headers, k)
	}

	t := table.NewWriter()
	t.AppendHeader(headers)

	for _, row := range d {
		var values []interface{}

		for _, column := range headers {
			values = append(values, row[column.(string)])
		}

		t.AppendRow(values)
	}

	fmt.Println(t.Render())
}
