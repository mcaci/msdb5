package assets

import (
	"html/template"

	"github.com/mcaci/ita-cards/set"
)

type Tmpl struct {
	n string
	f map[string]interface{}
	t string
}

func New(n string, f map[string]interface{}, t string) *Tmpl {
	return &Tmpl{n: n, f: f, t: t}
}

func (t *Tmpl) ToTmpl() (*template.Template, error) {
	return template.New(t.n).Funcs(t.f).Parse(t.t)
}

func Hand(pl interface{ Hand() *set.Cards }) func() (*template.Template, error) {
	return New("hand", map[string]interface{}{"hand": pl.Hand}, `{{ print "Hand"}}{{ range $i, $el:= hand }}<div>{{printf "(%d) %s" $i $el}}</div>{{ end }}<br/>`).ToTmpl
}

func Label(l string) func() (*template.Template, error) {
	return New("label", nil, l+`<div>{{printf "%s" .Label}}</div><br/>`).ToTmpl
}
