# PocketGoat Development Guide for AI Agents

## Build/Lint/Test Commands
```bash
mise dev                 # Run dev server with hot reload (air)
mise build              # Build production binary
mise templ              # Generate Go code from .templ files
mise css                # Generate CSS from Tailwind
go mod tidy             # Install/update Go dependencies
go test ./...           # Run all tests
go vet ./...            # Run Go linter
```

## Code Style Guidelines
- **Imports**: Group by stdlib, external deps, then local packages (pocketgoat/)
- **Templates**: Use .templ files with @layouts.BaseLayout wrapper for pages
- **Icons**: Copy SVGs directly into templates/icons/*.templ from lucide.dev/svgl.app
- **Routes**: Define in main.go OnServe() - specific routes before wildcards
- **Static**: Serve from pb_public/ via /static/{path...} route
- **Alpine.js**: Use x-data for interactivity, extend base layout data
- **Styling**: Tailwind CSS + DaisyUI components only
- **Error Handling**: Return proper HTTP status codes in handlers
- **PocketBase**: Use core.RequestEvent for request handling