package input

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
	words := strings.Split(d.Text, " ")
	currentWord := words[len(words)-1]

	s := []prompt.Suggest{}

	if len(strings.Trim(currentWord, "")) == 0 {
		return s
	}

	for _, keyword := range keywords {
		if strings.Contains(keyword, strings.ToUpper(currentWord)) {
			s = append(s, prompt.Suggest{
				Text: keyword,
			})
		}
	}

	return s
}
