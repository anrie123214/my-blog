# Repository Guidelines

## Project Structure & Module Organization

This repository is a small Go web application:

- `main.go` contains the HTTP server, routes, and template loading.
- `templates/` contains HTML pages, including `index.html` and `about.html`.
- `static/css/` contains stylesheets; `static/js/` contains browser scripts; and `static/images/` contains image assets.
- `go.mod` defines the Go module and language version.

There is currently no dedicated test directory. Add Go tests beside the code they cover, using files named `*_test.go`.

## Build, Test, and Development Commands

Run these commands from the repository root:

```bash
go run .       # Start the server on http://localhost:8080
go build .     # Compile the application
go test ./...  # Run all Go tests
```

Use `gofmt -w *.go` after changing Go source. Keep templates and static assets formatted consistently with their surrounding files.

## Coding Style & Naming Conventions

Follow standard Go conventions: tabs and `gofmt` for Go code, mixed-case exported names, and concise lower-case names for local variables and handlers. Keep HTTP routes and template names explicit and readable. Use semantic HTML, lower-case kebab-style CSS class names, and descriptive asset filenames.

## Testing Guidelines

Use Go’s standard `testing` package. Name tests `TestThing` and keep them focused on route behavior, template rendering, or other observable behavior. Run `go test ./...` before submitting changes. Add coverage for new server logic when practical.

## Commit & Pull Request Guidelines

Existing commits use short, informal descriptions. New commits should use a concise imperative subject, for example, `Improve header layout`. Keep unrelated changes separate.

Pull requests should explain the change, list verification commands and results, and include screenshots for visible HTML/CSS changes. Mention any route, configuration, or deployment considerations.

## Security & Configuration Tips

Do not commit credentials or local configuration. `.env` files are ignored; use environment variables for secrets and provide safe example values through documentation when needed.

## Agent Change Approval

Before modifying any repository file, ask the user for explicit approval. Read-only inspection is allowed, but do not edit, stage, commit, or push changes without confirmation. Keep unrelated working-tree changes untouched.
