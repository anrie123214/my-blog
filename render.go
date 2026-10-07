package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type pageData struct {
	Markdown template.HTML
	Category string
	Date     string
	Author   string
}

func renderMarkdown(path string) (template.HTML, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	p := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs)
	renderer := mdhtml.NewRenderer(mdhtml.RendererOptions{
		Flags: mdhtml.CommonFlags,
	})

	return template.HTML(markdown.ToHTML(source, p, renderer)), nil
}

func renderTemplate(w http.ResponseWriter, r *http.Request, page string, data any) error {
	pageFile := "templates/" + page + ".html" //concat string file name and dir
	templateFiles := []string{pageFile}
	templateName := "content"

	if r.Header.Get("HX-Request") != "true" {
		templateFiles = append([]string{"templates/index.html"}, pageFile)
		templateName = "layout"
	}

	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return fmt.Errorf("parse templates for %q: %w", page, err)
	}

	if err := tmpl.ExecuteTemplate(w, templateName, data); err != nil {
		return fmt.Errorf("execute %q template for %q: %w", templateName, page, err)
	}

	return nil
}

func newMux() http.Handler {
	mux := http.NewServeMux()
	staticFiles := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := renderMarkdown("docs/blogs/NamespacesAndCgroups.md")
		if err != nil {
			http.Error(w, "Could not render Markdown", http.StatusInternalServerError)
			return
		}
		if err := renderTemplate(w, r, "home", pageData{
			Markdown: content,
			Category: "Containers",
			Date:     "October 7, 2026",
			Author:   "Nyatt",
		}); err != nil {
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
