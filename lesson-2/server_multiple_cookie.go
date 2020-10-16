package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", firstHandle)
	router.HandleFunc("/user", helloUsername)

	log.Fatal(http.ListenAndServe(":8091", router))
}

func firstHandle(w http.ResponseWriter, req *http.Request) {
	cookie := &http.Cookie{
		Name:  "firstCookie",
		Value: "firstCookieValue",
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("Привет, мир!"))
}

func helloUsername(wr http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("firstCookie")
	if err != nil {
		fmt.Fprintf(wr, "Привет, %s!", req.URL.Query().Get("name"))
		return
	}
	fmt.Fprintf(wr, "Привет, %s! Я вижу вы побывали на / странице=)", req.URL.Query().Get("name"))
}
