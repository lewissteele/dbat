package input

import (
	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
    {Text: "select"},
    {Text: "from"},
  }
}
