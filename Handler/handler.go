package handler

import (
	asciiartweb "asciiartweb/Tools"
	"fmt"
	"html/template"
	"net/http"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) > 200 {
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

	fmt.Fprintln(w, result)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("Template/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 File not existing", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 File not existing", http.StatusInternalServerError)
		return
	}
}
