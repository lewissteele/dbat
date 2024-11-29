package input

import (
	"sort"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/c-bata/go-prompt"
	"github.com/lewissteele/dbat/internal/db"
)

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	if len(strings.TrimSpace(d.Text)) == 0 {
		return s
	}

	words := strings.Split(d.Text, " ")
	currentWord := words[len(words)-1]

	if len(strings.TrimSpace(currentWord)) == 0 {
		return s
	}

	s = append(s, similarity(currentWord, keywords[:])...)
	s = append(s, similarity(currentWord, db.Tables)...)
	s = append(s, similarity(currentWord, db.Databases)...)

	return s
}

func similarity(needle string, haystack []string) []prompt.Suggest {
	m := make(map[float64]string)
	s := []prompt.Suggest{}
	metric := metrics.JaroWinkler{}

	for _, v := range haystack {
		if needle == v {
			return s
		}

		if strings.Contains(v, needle) {
			similarity := strutil.Similarity(
				needle,
				v,
				&metric,
			)

			m[similarity] = v
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
