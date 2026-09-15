package handler

import (
	asciiartweb "asciiartweb/Tools"
	"html/template"
	"net/http"
)

const (
	maxlength = 200
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) > maxlength {
		http.Error(w, "400 Invalid text length", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(text); i++ {
		if (text[i] < 32 || text[i] > 126) && text[i] != 10 && text[i] != 13 {
			http.Error(w, "400 Invalid character in text", http.StatusBadRequest)
			return
		}
	}

	lines, err := asciiartweb.Loadbanner(banner)
	if err != nil {
		http.Error(w, "404 Banner not found", http.StatusNotFound)
		return
	}

	result := asciiartweb.AsciiPrint(text, lines)

	if err := RenderIndex(w, result); err != nil {
		http.Error(w, "500 File not existing", http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := RenderIndex(w, ""); err != nil {
		http.Error(w, "500 File not existing", http.StatusInternalServerError)
	}
}

func RenderIndex(w http.ResponseWriter, data string) error {
	tmpl, err := template.ParseFiles("Template/index.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}
