package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
