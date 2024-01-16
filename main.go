package main

import (
	"fmt"
	"gocr/handlers"
	"gocr/utils"
	"log"
	"net/http"
)

func main() {
	conf := &utils.Config
	utils.LoadConfig()

	http.HandleFunc("/", handlers.UploadHandler)

	fmt.Println("HTTP server is listening on " + conf.ListenAddr)
	err := http.ListenAndServe(conf.ListenAddr, nil)
	log.Fatalln(err)
}
