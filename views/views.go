package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout string

}

func addTemplateExt (files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}


// addTemplatePath takes in a slice of strings representing file paths
// for templates, and it prepends the TemplateDir directory to each string in the slice
func addTemplatePath (files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}


func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View {
		Template: t,
		Layout: layout,
	}
}

func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}
