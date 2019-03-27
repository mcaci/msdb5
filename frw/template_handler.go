package frw

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateHandler templ represents a single template
type TemplateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// NewTemplateHandler func
func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{filename: "msdb5.html"}
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("frw/templates",
			t.filename)))
	})
	t.templ.Execute(w, r)
}
