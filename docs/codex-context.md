# Codex Project Context

Last updated: 2026-10-07

## Current state

This is a small Go web application served on `http://localhost:8080`.

- `main.go` contains the entry point and calls `RenderAndServe()`.
- `render.go` contains the HTTP server, routes, and template rendering.
- `templates/index.html` is the shared full-page layout named `layout`.
- `templates/home.html`, `templates/about.html`, and `templates/posts.html` define page content named `content`.
- `static/css/style.css` contains the layout, responsive navigation drawer, theme colors, and button styles.
- `static/js/navbutton.js` controls the hamburger menu and light/dark mode.
- `static/js/htmx.js` is the checked-in HTMX client library.
- `static/images/logo.png` is the site logo.

There are currently no Go test files. The working tree was clean after the last push.

## Routes and rendering

Routes currently available:

- `/` renders `home.html`.
- `/about` renders `about.html`.
- `/posts` renders `posts.html`.
- `/static/` serves files from `static/`.

`renderTemplate` checks the `HX-Request` header:

- Normal browser requests parse `templates/index.html` plus the requested page and execute the `layout` template.
- HTMX requests parse only the requested page and execute its `content` template.

This allows `/about` and `/posts` to work both through HTMX navigation and after a full browser reload. `hx-push-url="true"` updates the address bar but does not itself provide full-page reload behavior; the server-side header check handles that.

## Navigation and theme behavior

The header contains:

- The logo link.
- A hamburger button for small screens.
- A right-side mobile navigation drawer using CSS transforms.
- Home, About, and Posts links.
- A theme toggle button.

The theme toggle:

- Adds or removes `dark-mode` on `<body>`.
- Stores the choice in `localStorage` under `theme`.
- Uses the moon SVG in light mode.
- Replaces it with the sun SVG in dark mode.
- Changes the label and accessible `aria-label` between “Dark mode” and “Light mode”.

The moon and sun SVGs are inline in `templates/index.html`/`static/js/navbutton.js`; there is no required external theme icon asset.

## CSS decisions

`style.css` is organized into sections for reset, theme colors, page base, header, logo, desktop navigation, controls, mobile drawer, and main content.

Important current values:

- Header padding: `16px 32px`.
- Desktop navigation gap: `12px`.
- Desktop navigation link padding: `24px`.
- Mobile drawer width: `min(80vw, 320px)`.
- Mobile drawer slides in from the right with `transform: translateX(...)`.

## Git and repository decisions

- The default branch is `main`.
- Remote: `origin` points to the project’s GitHub repository.
- GoLand `.idea/` files were removed from Git tracking and remain ignored locally.
- The repository instruction says that an explicit commit/push request uses `git add .`, followed by a staged-diff review for unrelated changes.
- Before modifying repository files, ask the user for explicit approval.
- Keep unrelated working-tree changes untouched.

Recent relevant commits:

- `dc88b2f` — Improve theme toggle styling
- `0eeb42d` — Stop tracking GoLand project files
- `9d7a10f` — Improve template rendering and mobile navigation

## Development workflow

Run from the repository root:

```bash
go run .
go build .
go test ./...
```

Templates are parsed on every request, which is convenient during development: HTML changes only require a browser refresh. GoLand’s static HTML preview does not run the Go routes; use the Go application at `http://localhost:8080` for route testing.

## Known cleanup items

- `header` uses `var(--header-border)`, but the current `:root` and `body.dark-mode` blocks do not define that variable. Add it before relying on the header border.
- `main { ... }` styles exist, but the layout currently uses `<div id="content">` rather than a `<main>` element.
- `http.ListenAndServe` errors are currently ignored.
- `renderTemplate` returns generic error messages (“Parsing not good” and “Layout's faulty”); these could be made more descriptive later.
- There is no automated route or template test coverage yet.
