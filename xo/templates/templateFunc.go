package templates

import (
	"html/template"
	"strings"

	"github.com/kenshaw/snaker"
)

var HelperFunc template.FuncMap = template.FuncMap{
	"camelCase": func(input string) string {
		return snaker.SnakeToCamel(strings.ReplaceAll(input, "/[^a-zA-Z0-9]/g", "_"))
	},
	"joinWith":  joinWith,
	"shortName": shortName,
}

func joinWith(with string, values ...string) string {
	return strings.Join(values, with)
}

// fetch only uppercase values
func shortName(name string) string {
	short := strings.ToLower(strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r
		}
		return -1
	}, name))

	if len(short) == 0 {
		short = strings.ToLower(name[0:1])
	}
	return short
}
