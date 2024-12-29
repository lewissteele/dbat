package table

import (
	"fmt"
	"github.com/gosuri/uitable"
)

func Render(d []map[string]interface{}) {
	t := uitable.New()
	t.MaxColWidth = 80
	t.Wrap = false

	var headers []interface{}
	for k, _ := range d[0] {
		headers = append(headers, k)
	}

	for _, row := range d {
		for _, header := range headers {
			k := header.(string)
			t.AddRow(k, row[k])
		}

		t.AddRow("")
	}

	fmt.Println(t)
}
