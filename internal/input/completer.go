package input

import (
	"sort"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
	m := make(map[float64]string)
	s := []prompt.Suggest{}

	if len(strings.TrimSpace(d.Text)) == 0 {
		return s
	}

	words := strings.Split(d.Text, " ")
	currentWord := words[len(words)-1]

	if len(strings.TrimSpace(currentWord)) == 0 {
		return s
	}

	metric := metrics.JaroWinkler{}

	for _, keyword := range keywords {
		if strings.Contains(keyword, currentWord) {
			similarity := strutil.Similarity(
				currentWord,
				keyword,
				&metric,
			)

			m[similarity] = keyword
		}
	}

	keys := make([]float64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	for _, k := range keys {
		s = append(s, prompt.Suggest{
			Text: m[k],
		})
	}

	return s
}
