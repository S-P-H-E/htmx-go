package handler

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed pages/*.html
var templatesFS embed.FS

var tmpl = template.Must(template.ParseFS(templatesFS, "pages/*.html"))

type Count struct {
	Count int
}

func Handler(w http.ResponseWriter, r *http.Request) {
	count := Count{Count: 0} // State is managed client-side with Alpine.js
	tmpl.ExecuteTemplate(w, "index", count)
}
