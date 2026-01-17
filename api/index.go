package handler

import (
	"embed"
	"html/template"
	"net/http"
	"sync"
)

//go:embed pages/*.html
var templatesFS embed.FS

type Count struct {
	Count int
}

var (
	tmpl     *template.Template
	tmplOnce sync.Once
)

func loadTemplate() {
	tmplOnce.Do(func() {
		tmpl, _ = template.ParseFS(templatesFS, "pages/*.html")
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
