package main

import "net/http"

func newMux() http.Handler {
	mux := http.NewServeMux()
	staticFiles := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		post, err := loadPost("docs/blogs/ContainersOnLinuxP1.md")
		if err != nil {
			http.Error(w, "Could not render Markdown", http.StatusInternalServerError)
			return
		}
		if err := renderTemplate(w, r, "home", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if err := renderTemplate(w, r, "about", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if err := renderTemplate(w, r, "posts", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return mux
}

func RenderAndServe() {
	http.ListenAndServe(":8080", newMux())
}
