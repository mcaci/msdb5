package frw

import (
	"html/template"
	"net/http"
	"os"
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
	path := templatePath(t.filename, func(testPath string) bool {
		_, err := os.Stat(filepath.Join(testPath, t.filename))
		return !os.IsNotExist(err)
	})
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join(path, t.filename)))
	})
	t.templ.Execute(w, r)
}

func templatePath(filename string, fileExists func(string) bool) string {
	const localPath = "frw/templates"
	localFile := filepath.Join(localPath, filename)
	if fileExists(localFile) {
		return localFile
	}
	const packagePath = "github.com/nikiforosFreespirit/msdb5/"
	return packagePath + localFile
}
