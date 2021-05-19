package assets

import (
	"html/template"
)

func List(k string, l interface{}) func() (*template.Template, error) {
	return func() (*template.Template, error) {
		return template.New("list").Funcs(map[string]interface{}{k: l}).Parse(k + `<div>{{ range $i, $el:=` + k + `}}<div>{{printf "(%d) %s" $i $el}}</div>{{ end }}</div><br/>`)
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
