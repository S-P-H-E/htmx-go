package main

import (
	"net/http"
	handler "github.com/S-P-H-E/htmx-go/api"
)

func main() {
	// Serve static files
	http.Handle("/favicon.ico", http.FileServer(http.Dir("public")))

	// Use the same Handler as Vercel
	http.HandleFunc("/", handler.Handler)

	http.ListenAndServe(":3000", nil)
}
