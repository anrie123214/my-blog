# Project Context

Last updated: 2026-10-08

## Conversation and workflow context

- This repository is intended to be worked on from both a laptop and a PC.
- Shared project knowledge should be synchronized through Git.
- `AGENTS.md` contains repository-specific instructions for Codex.
- This file contains shared application architecture, decisions, roadmap notes, and relevant workflow context.
- Machine-specific Codex settings, personal preferences, credentials, and local paths should remain outside the repository.
- `TASK.md` contains the short project goal and requirements.
- Before editing, renaming, deleting, staging, committing, or pushing repository files, obtain explicit user approval. That approval was given for the current documentation organization and push.

## Current architecture

This is a small Go web app served at `http://localhost:8080`.

- `main.go` starts the server.
- `render.go` defines HTTP routes and server startup.
- `content.go` parses post front matter and loads post data.
- `markdown.go` converts Markdown to HTML with `gomarkdown`.
- `templates.go` renders full layouts or HTMX content fragments.
- `templates/index.html` is the shared layout and keeps the header/navigation mounted.
- `templates/home.html`, `about.html`, and `posts.html` provide swappable page content.
- `static/css/style.css` contains the responsive layout and article styling.
- `static/js/navbutton.js` handles the menu and theme toggle; HTMX handles content navigation and scroll reset declaratively in the shared layout.

## Routes and HTMX

- `/` renders `docs/blogs/ContainersOnLinuxP1.md` as the current featured Markdown article.
- `/about` renders the personal profile and freelance-oriented About page.
- `/posts` is planned as a table-of-contents-style article list.
- `/static/` serves static assets.

Normal requests receive the full layout. Requests with `HX-Request: true` receive only the page content, which replaces `main#content`. Navigation uses `hx-push-url` and an HTMX `after:swap` hook so direct links, browser refreshes, and scroll position continue to work.

The `.site-header` is sticky and hides while scrolling down, then returns while scrolling up. The mobile drawer locks page scrolling while open, includes a white close button, and closes when a link or outside area is clicked.

## Content direction

- Home currently recaps Part 1 of the Containers on Linux series, inspired by the Podman blog format.
- Article headers can show category, date, and author pills.
- About should promote freelance work, LinkedIn, and related links.
- Posts should list the Markdown articles in `docs/blogs/`, with individual article routes added later.
- Current articles include `ContainersOnLinuxP1.md` and `ContainersOnLinuxP2.md`.
- The final visual design may change when a reference design is provided.

## Development decisions

- Keep Go files at the repository root while the app remains small; do not add `src/` yet.
- Store articles in `docs/blogs/` with YAML-like front matter for title, date, category, author, and slug, then render them with `gomarkdown/markdown`.
- Keep post loading and Markdown conversion separate from route and template concerns.
- Keep JavaScript minimal and use HTMX for page navigation.
- Keep `.article-header` separate from the sticky `.site-header`.
- Go tests live beside the code in `*_test.go` files.
- Run `gofmt`, `go build .`, and `go test ./...` after Go changes.
- Ask for approval before editing, renaming, deleting, staging, or committing files.

## Next steps

- Turn `/posts` into a generated article list using post front matter.
- Add front matter to Part 2 and validate all posts during loading.
- Add individual article routes using post slugs.
- Add freelance and LinkedIn links to the About page.
- Consider improving `http.ListenAndServe` error handling.
- Update `document.title` when the page changes.
