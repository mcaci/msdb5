package frw

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/text/language"
)

// TemplateHandler templ represents a single template
type TemplateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// NewTemplateHandler func
func NewTemplateHandler(lang language.Tag) *TemplateHandler {
	var tmpl string
	switch lang {
	case language.Italian:
		tmpl = "msdb5-it.html"
	default:
		tmpl = "msdb5.html"
	}
	return &TemplateHandler{filename: tmpl}
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := templatePath(t.filename, func(testPath string) bool {
		_, err := os.Stat(filepath.Join(testPath, t.filename))
		return !os.IsNotExist(err)
	})
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(path))
	})
	t.templ.Execute(w, r)
}

func templatePath(filename string, fileExists func(string) bool) string {
	localFile := filepath.Join("frw/templates", filename)
	if fileExists(localFile) {
		return localFile
	}
	return "github.com/mcaci/msdb5/" + localFile
}
