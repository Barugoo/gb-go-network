package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/{id}/posts", firstHandle)
	router.HandleFunc("/other", secondHandle)

	log.Println("Starting server at :8092")
	log.Fatal(http.ListenAndServe(":8092", router))
}

func firstHandle(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("Привет, мир!"))
}

func secondHandle(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("Привет, мир! (secondHandle)"))
}
