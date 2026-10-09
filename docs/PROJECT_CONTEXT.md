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

- `/` renders the home landing page with the generated blog collection as featured posts.
- `/about` renders the personal profile and freelance-oriented About page.
- `/posts` renders a generated article list from the Markdown collection.
- `/posts/{slug}` renders an individual Markdown article.
- `/static/` serves static assets.

Normal requests receive the full layout. Requests with `HX-Request: true` receive only the page content, which replaces `main#content`. Navigation uses `hx-push-url` and an HTMX `after:swap` hook so direct links, browser refreshes, and scroll position continue to work.

The `.site-header` is sticky and hides while scrolling down, then returns while scrolling up. The mobile drawer locks page scrolling while open, includes a white close button, and closes when a link or outside area is clicked.

## Content direction

- Home has a centered welcome hero, a smooth scroll link, and generated featured post cards with a Read More link to the posts archive.
- Article headers can show category, date, and author pills.
- About should promote freelance work, LinkedIn, and related links.
- Posts are loaded from `docs/blogs/`, sorted by date, and linked through their front matter slugs.
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

- Add descriptive browser tab titles for the home, posts, article, and about pages.
- Add freelance and LinkedIn links to the About page.
- Consider improving `http.ListenAndServe` error handling.

## Work log — 2026-10-09

- Added older/newer navigation to article pages using server-selected neighboring posts.
- Added HTMX navigation attributes and responsive styling for article navigation buttons.
- Applied `show:top` to all HTMX page and article links so swapped content returns to the top of the target content area.
- Moved the home welcome section toward the top of the page instead of vertically centering it.
- Kept F5 as a full browser reload and added reload-specific scroll reset behavior so refreshed pages start at the top.
- Added browser `prefers-color-scheme` detection for the initial theme, with explicit `localStorage` choices taking priority and system theme changes applied when no manual choice exists.
- Current verification note: `gofmt` and `git diff --check` pass. Go build/test verification may require writable Go caches and network access for the Markdown dependency in the development environment.
