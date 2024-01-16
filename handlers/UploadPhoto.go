package handlers

import (
	"fmt"
	"gocr/utils"
	"io"
	"net/http"
	"os"
)

const uploadDir = "./uploads/"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Create upload directory
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil && err != os.ErrExist {
		fmt.Println("Error creating upload directory:", err)
		return
	}

	// Parse form
	err = r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Receive file
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := fmt.Sprintf("%s%s", uploadDir, handler.Filename)
	newFile, _ := os.Create(filePath)

	defer newFile.Close()
	io.Copy(newFile, file)

	text, err := utils.OCR(filePath)

	if err == nil {
		fmt.Fprintf(w, "\n%s\n\n", text)
	} else {
		fmt.Fprintf(w, "\nFailed getting text content!\n\n")
	}
}
