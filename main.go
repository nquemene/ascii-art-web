package main

import (
	handler "asciiartweb/Handler"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("GET /Style/", http.StripPrefix("/Style/", http.FileServer(http.Dir("Style"))))
	http.HandleFunc("GET /{$}", handler.IndexHandler)
	http.HandleFunc("POST /ascii-art", handler.AsciiArtHandler)

	fmt.Println("http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("500 Internal Server Error")
		return
	}
}
