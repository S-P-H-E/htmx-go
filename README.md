# Go HTMX Counter

A simple counter web application built with Go, Echo, and HTMX.

## Tech Stack

- **Go** - Backend
- **Echo** - Web framework
- **HTMX** - Frontend interactivity without JavaScript
- **Tailwind CSS** - Styling

## Getting Started

```bash
# Install dependencies
go mod download

# Run the server
go run main.go
```

The app runs at `http://localhost:3000`.

## How It Works

Click the button to increment the counter. HTMX handles the POST request and swaps the updated count into the page without a full reload.
