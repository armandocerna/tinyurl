package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"tinyurl/tinyurl"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/new", tinyurl.NewTinyUrlHandler).Methods("GET")
	r.HandleFunc("/new", tinyurl.NewTinyUrlPostHandler).Methods("POST")
	r.HandleFunc("/{tinyUrl}", tinyurl.TinyUrlRedirectHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
