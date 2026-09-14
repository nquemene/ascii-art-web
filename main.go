package main

import (
	handler "asciiartweb/Handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /{$}", handler.IndexHandler)
	http.HandleFunc("POST /ascii-art", handler.AsciiArtHandler)

	fmt.Println("http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error 500 : Internal Server Error")
		return
	}
}
