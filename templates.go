package main

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type views struct {
	templates map[string]*template.Template
	root      []string
}

func (v *views) Register(name string) {
	v.templates[name] = template.Must(template.New("").ParseFiles(
		filepath.Join(append(v.root, "layout.html")...),
		filepath.Join(append(v.root, name+".html")...),
	))
}

func (v *views) Render(w io.Writer, name string, data interface{}) {
	if tmpl, exists := v.templates[name]; exists {
		tmpl.ExecuteTemplate(w, "base", data)
		return
	}
	fmt.Fprintf(w, "Could not find template '%s'", name)
}

func loadViews(root string) *views {
	return &views{
		templates: make(map[string]*template.Template),
		root:      strings.Split(root, "/"),
	}
}
