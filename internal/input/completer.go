package input

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	if len(strings.TrimSpace(d.Text)) == 0 {
		return s
	}

	words := strings.Split(d.Text, " ")
	currentWord := words[len(words)-1]

	for _, keyword := range keywords {
		if strings.Contains(keyword, currentWord) {
			s = append(s, prompt.Suggest{
				Text: keyword,
			})
		}
	}

	return s
}
