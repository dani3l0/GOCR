package main

import (
	"gocr/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.UploadHandler)
	http.ListenAndServe(":8080", nil)
}
