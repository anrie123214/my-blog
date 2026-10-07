# Codex Project Context

Last updated: 2026-10-07

## Current architecture

This is a small Go web app served at `http://localhost:8080`.

- `main.go` starts the server.
- `render.go` defines routes, template rendering, and Markdown conversion with `gomarkdown`.
- `templates/index.html` is the shared layout and keeps the header/navigation mounted.
- `templates/home.html`, `about.html`, and `posts.html` provide swappable page content.
- `static/css/style.css` contains the responsive layout and article styling.
- `static/js/navbutton.js` handles the menu and theme toggle.

## Routes and HTMX

- `/` renders the current featured Markdown article.
- `/about` is planned as a personal name card and freelance profile.
- `/posts` is planned as a table-of-contents-style article list.
- `/static/` serves static assets.

Normal requests receive the full layout. Requests with `HX-Request: true` receive only the page content, which replaces `main#content`. Navigation uses `hx-push-url` so direct links and browser refreshes continue to work.

## Content direction

- Home should recap the newest article, inspired by the Podman blog format.
- Article headers can show category, date, and author pills.
- About should promote freelance work, LinkedIn, and related links.
- Posts should list Markdown articles, with individual article routes added later.
- The final visual design may change when a reference design is provided.

## Development decisions

- Keep Go files at the repository root while the app remains small; do not add `src/` yet.
- Store articles in `docs/blogs/` and render them with `gomarkdown/markdown`.
- Keep JavaScript minimal and use HTMX for page navigation.
- Go tests live beside the code in `*_test.go` files.
- Run `gofmt`, `go build .`, and `go test ./...` after Go changes.
- Ask for approval before editing, renaming, deleting, staging, or committing files. Read-only commands are okay without approval.

## Next steps

- Turn `/posts` into a generated article list.
- Add metadata handling for article title, date, author, and category.
- Build the About name card and freelance links.
- Consider improving `http.ListenAndServe` error handling.
- Update `document.title` when the page changes.
