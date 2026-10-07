package main

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, r *http.Request, page string) {
	pageFile := "templates/" + page + ".html" //concat string file name and dir
	templateFiles := []string{pageFile}
	templateName := "content"

	if r.Header.Get("HX-Request") != "true" {
		templateFiles = append([]string{"templates/index.html"}, pageFile)
		templateName = "layout"
	}

	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, "Parsing not good", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, templateName, nil); err != nil {
		http.Error(w, "Layout's faulty", http.StatusInternalServerError)
	}
}

func RenderAndServe() {
	mux := http.NewServeMux()
	staticFiles := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "home")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "about")
	})
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "posts")
	})
	http.ListenAndServe(":8080", mux)
}
