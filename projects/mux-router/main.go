package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getBookDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Book details will be here")
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "Title: %s\n", title)
	fmt.Fprintf(w, "Page: %s\n", page)
}

func main() {
	fmt.Println("HTTP Server main function.")
	r := mux.NewRouter()
	fmt.Println("Router initialized:", r)
	r.HandleFunc("/books/{title}/page/{page}", getBookDetail)

	userrouter := r.PathPrefix("/users").Subrouter()
	userrouter.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Info will be here")
	})
	userrouter.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Settings will be here")
	})

	fmt.Println("Starting server at port 4000")

	http.ListenAndServe(":4000", r)

}
