package table

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"
)

func Render(data []map[string]interface{}) {
	var columns []string
	for key, _ := range data[0] {
		columns = append(columns, key)
	}

	if len(columns) == 1 {
		for _, row := range data {
			for _, column := range columns {
				fmt.Println(row[column])
			}
		}
		return
	}

	sort.Slice(columns, func(a, b int) bool {
		return len(columns[a]) < len(columns[b])
	})

	padding := 0
	for _, s := range columns {
		l := len(s)
		if l > padding {
			padding = l
		}
	}

	for idx, row := range data {
		for _, column := range columns {
			if len(columns) == 1 {
				fmt.Println(row[column])
				continue
			}

			color.Set(color.FgGreen)
			fmt.Printf("%*s", padding, column)

			color.Set(color.Reset)
			fmt.Println("", row[column])
		}

		if len(data)-1 != idx {
			fmt.Println(strings.Repeat("-", 80))
		}
	}
}
