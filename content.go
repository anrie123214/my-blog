package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type pageData struct {
	Title    string
	Markdown template.HTML
	Category string
	Date     string
	Author   string
	Slug     string
}

type articlePageData struct {
	pageData
	Older *pageData
	Newer *pageData
}

type postMetadata struct {
	Title    string
	Date     string
	Category string
	Author   string
	Slug     string
}

func parseFrontMatter(source []byte) (postMetadata, []byte, error) {
	lines := strings.Split(string(source), "\n")
	if len(lines) == 0 || strings.TrimSpace(strings.TrimSuffix(lines[0], "\r")) != "---" {
		return postMetadata{}, source, nil
	}

	closingLine := -1
	for index := 1; index < len(lines); index++ {
		line := strings.TrimSpace(strings.TrimSuffix(lines[index], "\r"))
		if line == "---" {
			closingLine = index
			break
		}
	}
	if closingLine == -1 {
		return postMetadata{}, nil, fmt.Errorf("front matter is missing a closing delimiter")
	}

	var metadata postMetadata
	for _, line := range lines[1:closingLine] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		key, value, ok := strings.Cut(line, ":")
		if !ok {
			return postMetadata{}, nil, fmt.Errorf("invalid front matter line %q", line)
		}

		value = strings.Trim(strings.TrimSpace(value), "\"'")
		switch strings.TrimSpace(key) {
		case "title":
			metadata.Title = value
		case "date":
			metadata.Date = value
		case "category":
			metadata.Category = value
		case "author":
			metadata.Author = value
		case "slug":
			metadata.Slug = value
		default:
			return postMetadata{}, nil, fmt.Errorf("unknown front matter field %q", strings.TrimSpace(key))
		}
	}

	body := strings.Join(lines[closingLine+1:], "\n")
	body = strings.TrimPrefix(body, "\n")
	return metadata, []byte(body), nil
}

func loadPost(path string) (pageData, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return pageData{}, err
	}

	metadata, body, err := parseFrontMatter(source)
	if err != nil {
		return pageData{}, err
	}
	if metadata.Title == "" || metadata.Date == "" || metadata.Category == "" || metadata.Author == "" || metadata.Slug == "" {
		return pageData{}, fmt.Errorf("post %q is missing required front matter", path)
	}

	return pageData{
		Title:    metadata.Title,
		Markdown: renderMarkdownSource(body),
		Category: metadata.Category,
		Date:     metadata.Date,
		Author:   metadata.Author,
		Slug:     metadata.Slug,
	}, nil
}

func loadPosts(directory string) ([]pageData, error) {
	paths, err := filepath.Glob(filepath.Join(directory, "*.md"))
	if err != nil {
		return nil, fmt.Errorf("find posts in %q: %w", directory, err)
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("no Markdown posts found in %q", directory)
	}

	posts := make([]pageData, 0, len(paths))
	slugs := make(map[string]string, len(paths))
	for _, path := range paths {
		post, err := loadPost(path)
		if err != nil {
			return nil, err
		}
		if previousPath, exists := slugs[post.Slug]; exists {
			return nil, fmt.Errorf("duplicate post slug %q in %q and %q", post.Slug, previousPath, path)
		}

		slugs[post.Slug] = path
		posts = append(posts, post)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		if posts[i].Date == posts[j].Date {
			return posts[i].Title < posts[j].Title
		}
		return posts[i].Date > posts[j].Date
	})

	return posts, nil
}
