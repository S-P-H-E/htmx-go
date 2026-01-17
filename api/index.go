package handler

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Count struct {
	Count int
}

var (
	tmpl     *template.Template
	tmplOnce sync.Once
)

func loadTemplate() {
	tmplOnce.Do(func() {
		// Try relative path first (works when running from project root)
		var err error
		tmpl, err = template.ParseGlob("pages/*.html")

		// If that fails, try absolute path from current working directory
		if err != nil || tmpl == nil {
			if cwd, cwdErr := os.Getwd(); cwdErr == nil {
				// Build glob pattern: use filepath.Join for directory, then append wildcard
				basePath := filepath.Join(cwd, "pages")
				pattern := basePath + string(os.PathSeparator) + "*.html"
				tmpl, _ = template.ParseGlob(pattern)
			}
		}
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	loadTemplate()

	if tmpl == nil {
		http.Error(w, "Error: template not loaded", http.StatusInternalServerError)
		return
	}

	// Initialize count to 0 (state is managed client-side with Alpine.js)
	count := Count{Count: 0}

	// Execute template
	err := tmpl.ExecuteTemplate(w, "index", count)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
