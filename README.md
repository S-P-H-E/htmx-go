# Go HTMX Counter

A simple counter web application built with Go, HTMX, and Alpine.js. Deployable to Vercel.

## Tech Stack

- **Go** - Backend with `net/http` and embedded templates
- **HTMX** - Frontend interactivity
- **Alpine.js** - Client-side state management
- **Tailwind CSS** - Styling

## Getting Started

```bash
go mod download
go run .
```

The app runs at `http://localhost:3000`.

## How It Works

Templates are embedded into the binary using Go's `embed` package, making deployment simple. Alpine.js manages the counter state client-side.
