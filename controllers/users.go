package controllers

import (
	"fmt"
	"github.com/natecowen/lenslocked.com/views"
	"net/http"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email string `schema: "email""`
	Password string `schema: "password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

// Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}

}

//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)

}
