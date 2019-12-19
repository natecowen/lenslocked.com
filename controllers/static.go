package controllers

import "github.com/natecowen/lenslocked.com/views"

type Static struct {
	Home *views.View
	Contact *views.View
	FAQ *views.View
}

// initialize our static controller and our views
func NewStatic() *Static {
	return &Static {
		Home: views.NewView(
			"bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		FAQ: views.NewView("bootstrap", "static/faq"),
	}
}


