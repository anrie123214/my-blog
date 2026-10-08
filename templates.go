package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, r *http.Request, page string, data any) error {
	pageFile := "templates/" + page + ".html"
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
