package main

import (
	"fmt"
	"github.com/natecowen/lenslocked.com/views"
	"net/http"

	"github.com/gorilla/mux"
)

//Globals
var _404 http.Handler = http.HandlerFunc(notFound)

//Gloabals - VIEWS
var contactView *views.View
var faqView *views.View
var homeView *views.View

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func must(err error){
	if err !=nil {
		panic(err)
	}
}



func notFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 - Content Not Found</h1><p>Something went wrong, the page you are looking for does not exist. Please try again.</p>")
}

func main() {
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	homeView = views.NewView("bootstrap", "views/home.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = _404
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}