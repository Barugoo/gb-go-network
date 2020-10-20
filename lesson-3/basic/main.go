package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8093", router))
}
