package main

import (
	"fmt"
	"net/http"
)

// type :struct
type Router struct{}

var contactresp string = "<h1>Contact Page</h1><p>To get in touch,email me at <a href=\"mailto:arajeet.bapi@gmail.com\"> Arajeet Pal </a> </p>"

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handleFunc(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Header)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my website <h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, contactresp)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, r.URL.Path)
}

func main() {
	var router Router
	// type: function (quite weird)
	// var rx http.HandlerFunc = pathHandler
	//	http.HandleFunc("/", pathHandler)
	//	http.HandleFunc("/home", handleFunc)
	//	http.HandleFunc("/contact", contactHandler)
	// http.Handle("/", router)
	fmt.Println("Starting server on port: 8080")
	http.ListenAndServe(":8080", router)
}
