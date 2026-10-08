package main

import (
	"net/http"
	"strings"
)

func newMux() http.Handler {
	mux := http.NewServeMux()
	staticFiles := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		posts, err := loadPosts("docs/blogs")
		if err != nil {
			http.Error(w, "Could not render Markdown", http.StatusInternalServerError)
			return
		}
		if err := renderTemplate(w, r, "home", posts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if err := renderTemplate(w, r, "about", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts, err := loadPosts("docs/blogs")
		if err != nil {
			http.Error(w, "Could not load posts", http.StatusInternalServerError)
			return
		}
		if err := renderTemplate(w, r, "posts", posts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/posts/")
		posts, err := loadPosts("docs/blogs")
		if err != nil {
			http.Error(w, "Could not load posts", http.StatusInternalServerError)
			return
		}

		for _, post := range posts {
			if post.Slug == slug {
				if err := renderTemplate(w, r, "article", post); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		}

		http.NotFound(w, r)
	})

	return mux
}

func RenderAndServe() {
	http.ListenAndServe(":8080", newMux())
}
