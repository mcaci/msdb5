package templates

type Page struct {
	Title string
	Body  []byte
}

// Read at https://golang.org/doc/articles/wiki/
func (p *Page) save() error {
	filename := p.Title + "txt"
	return io.WriteFile(filename)
}
