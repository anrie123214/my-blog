package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	mux := http.NewServeMux()
	staticFiles := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "about.html", nil)
	})
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "posts.html", nil)
	})
	http.ListenAndServe(":8080", mux)
}
