package tinyurl

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var UrlDB = make(map[string]string)

func NewTinyUrlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/new.gohtml")
}

func NewTinyUrlPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")
	tu := CreateRandomString()
	UrlDB[tu] = url

	fmt.Fprintf(w, "\nURL: %s\nTiny: %s\n", url, tu)
	w.WriteHeader(http.StatusCreated)
	defer r.Body.Close()
}

func TinyUrlRedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if url, ok := UrlDB[vars["tinyUrl"]]; !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		http.Redirect(w, r, url, 302)
	}
	defer r.Body.Close()
}
