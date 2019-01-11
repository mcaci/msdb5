package main

import (
	"html/template"
	"log"
	"net/http"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

// Any struct
type Any struct {
	Title string
	Items []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Start func
func Start(w http.ResponseWriter, r *http.Request) {

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	writeOut(t, w, Any{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		}})

	writeOut(t, w, Any{
		Title: "My another page",
		Items: []string{}})
}

func writeOut(t *template.Template, w http.ResponseWriter, any Any) {
	err := t.Execute(w, any)
	check(err)
}
