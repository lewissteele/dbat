package input

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var keywords = []string{
	"from",
	"select",
	"update",
	"where",
}

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	if len(strings.Trim(d.Text, "")) == 0 {
		return s
	}

	for _, keyword := range keywords {
		if strings.Contains(keyword, d.Text) {
			s = append(s, prompt.Suggest{
				Text: keyword,
			})
		}
	}

	return s
}
