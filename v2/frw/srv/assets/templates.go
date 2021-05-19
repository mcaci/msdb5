package assets

import (
	"html/template"

	"github.com/mcaci/ita-cards/set"
)

func Hand(pl interface{ Hand() *set.Cards }) func() (*template.Template, error) {
	return func() (*template.Template, error) {
		return template.New("hand").Funcs(map[string]interface{}{"hand": pl.Hand}).Parse(`{{ print "Hand"}}{{ range $i, $el:= hand }}<div>{{printf "(%d) %s" $i $el}}</div>{{ end }}<br/>`)
	}
}
func Label(l string) func() (*template.Template, error) {
	return func() (*template.Template, error) {
		return template.New("label").Parse(l + `<div>{{printf "%s" .Label}}</div><br/>`)
	}
}
func Game() (*template.Template, error) {
	return template.ParseFiles("assets/game.html")
}
