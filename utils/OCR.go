package utils

import (
	"github.com/otiai10/gosseract/v2"
)

func OCR(imagePath string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(imagePath)
	client.SetLanguage(client.Languages...)

	return client.Text()
}
