package main

import (
	"fmt"
	"github.com/natecowen/lenslocked.com/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

//Globals
var _404 http.Handler = http.HandlerFunc(notFound)





func must(err error) {
	if err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 - Content Not Found</h1><p>Something went wrong, the page you are looking for does not exist. Please try again.</p>")
}

func main() {

	//controller import
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	//router functionality
	r := mux.NewRouter()
	r.NotFoundHandler = _404
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/faq", staticC.FAQ).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
