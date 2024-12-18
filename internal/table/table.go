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

		for _, header := range headers {
			k := header.(string)

			if s, ok := row[k].(string); ok {
				row[k] = ellipsis(s, 50)
			}

			values = append(values, row[k])
		}

		t.AppendRow(values)
	}

	fmt.Println(t.Render())
}

func ellipsis(s string, maxLen int) string {
    runes := []rune(s)

    if len(runes) <= maxLen {
        return s
    }

    if maxLen < 3 {
        maxLen = 3
    }

    return string(runes[0:maxLen-3]) + "..."
}
