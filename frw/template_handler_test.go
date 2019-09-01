package frw

import "testing"

func TestTemplatePath(t *testing.T) {
	s := templatePath("template.html", func(string) bool { return true })
	if s != "frw/templates/template.html" {
		t.Fatalf("Unexpected %s path", s)
	}
}

func TestTemplatePackagePath(t *testing.T) {
	s := templatePath("template.html", func(string) bool { return false })
	if s != "github.com/mcaci/msdb5/frw/templates/template.html" {
		t.Fatalf("Unexpected %s path", s)
	}
}
