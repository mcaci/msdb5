package assets

import (
	"html/template"
	"net/http"
)

func MustExecute(
	fTmpl func() (*template.Template, error),
	w http.ResponseWriter,
	data interface{},
) {
	tmpl, err := fTmpl()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
