package main

import (
	"html/template"
	"os"

	"github.com/gomarkdown/markdown"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func renderMarkdownSource(source []byte) template.HTML {
	p := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs)
	renderer := mdhtml.NewRenderer(mdhtml.RendererOptions{
		Flags: mdhtml.CommonFlags,
	})

	return template.HTML(markdown.ToHTML(source, p, renderer))
}

func renderMarkdown(path string) (template.HTML, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	_, body, err := parseFrontMatter(source)
	if err != nil {
		return "", err
	}

	return renderMarkdownSource(body), nil
}
