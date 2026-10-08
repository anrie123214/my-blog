package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	path := filepath.Join(t.TempDir(), "post.md")
	if err := os.WriteFile(path, []byte("# Test heading\n\nA paragraph."), 0o600); err != nil {
		t.Fatalf("write test Markdown: %v", err)
	}

	got, err := renderMarkdown(path)
	if err != nil {
		t.Fatalf("renderMarkdown returned an error: %v", err)
	}

	if !strings.Contains(string(got), "<h1 id=\"test-heading\">Test heading</h1>") {
		t.Errorf("rendered Markdown did not contain the expected heading: %s", got)
	}
	if !strings.Contains(string(got), "<p>A paragraph.</p>") {
		t.Errorf("rendered Markdown did not contain the expected paragraph: %s", got)
	}
}

func TestParseFrontMatter(t *testing.T) {
	source := []byte("---\ntitle: Test post\ndate: 2026-10-08\ncategory: Testing\nauthor: Nyatt\nslug: test-post\n---\n\n# Test heading")

	metadata, body, err := parseFrontMatter(source)
	if err != nil {
		t.Fatalf("parseFrontMatter returned an error: %v", err)
	}

	if metadata.Title != "Test post" || metadata.Date != "2026-10-08" || metadata.Category != "Testing" || metadata.Author != "Nyatt" || metadata.Slug != "test-post" {
		t.Fatalf("unexpected metadata: %+v", metadata)
	}
	if string(body) != "# Test heading" {
		t.Fatalf("unexpected body: %q", body)
	}
}

func TestRenderMarkdownStripsFrontMatter(t *testing.T) {
	path := filepath.Join(t.TempDir(), "post.md")
	source := []byte("---\ntitle: Hidden metadata\ndate: 2026-10-08\ncategory: Testing\nauthor: Nyatt\nslug: test-post\n---\n\n# Visible heading")
	if err := os.WriteFile(path, source, 0o600); err != nil {
		t.Fatalf("write Markdown: %v", err)
	}

	got, err := renderMarkdown(path)
	if err != nil {
		t.Fatalf("renderMarkdown returned an error: %v", err)
	}

	if strings.Contains(string(got), "Hidden metadata") {
		t.Errorf("rendered Markdown contained front matter: %s", got)
	}
	if !strings.Contains(string(got), "Visible heading") {
		t.Errorf("rendered Markdown did not contain the article body: %s", got)
	}
}

func TestHomeRendersFullLayout(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	newMux().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, recorder.Code, recorder.Body)
	}

	body := recorder.Body.String()
	for _, want := range []string{
		"<!DOCTYPE html>",
		"<main id=\"content\">",
		"class=\"article\"",
		"Containers on Linux",
		`class="article-category">Containers`,
		`class="article-meta">2026-10-07 //// Nyatt`,
	} {
		if !strings.Contains(body, want) {
			t.Errorf("full page did not contain %q", want)
		}
	}
}

func TestHomeHTMXRendersContentOnly(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("HX-Request", "true")

	newMux().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, recorder.Code, recorder.Body)
	}

	body := recorder.Body.String()
	if strings.Contains(body, "<!DOCTYPE html>") {
		t.Error("HTMX response unexpectedly contained the full page layout")
	}
	for _, want := range []string{"class=\"article\"", "Containers on Linux"} {
		if !strings.Contains(body, want) {
			t.Errorf("HTMX response did not contain %q", want)
		}
	}
}

func TestStaticPagesRenderFullLayout(t *testing.T) {
	for _, page := range []string{"/about", "/posts"} {
		t.Run(page, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, page, nil)

			newMux().ServeHTTP(recorder, request)

			if recorder.Code != http.StatusOK {
				t.Fatalf("expected status %d, got %d: %s", http.StatusOK, recorder.Code, recorder.Body)
			}
			if !strings.Contains(recorder.Body.String(), "<!DOCTYPE html>") {
				t.Error("response did not contain the full page layout")
			}
		})
	}
}

func TestRenderTemplateReportsParseErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/missing", nil)

	err := renderTemplate(recorder, request, "missing", nil)
	if err == nil {
		t.Fatal("renderTemplate did not report an error for a missing template")
	}
	if !strings.Contains(err.Error(), "parse templates") {
		t.Errorf("expected a parse error, got %v", err)
	}
}
