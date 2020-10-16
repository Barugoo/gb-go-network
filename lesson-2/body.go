package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	FirstName    string   `json:"firstName" xml:"firstName"`
	LastName     string   `json:"lastName" xml:"lastName"`
	PhoneNumbers []string `json:"phoneNumbers" xml:"phoneNumbers"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/person", personHandler)

	log.Fatal(http.ListenAndServe(":8093", router))
}

func personHandler(wr http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		wr.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	contentTypeHeader := req.Header.Get("Content-Type")

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	person := &Person{}
	switch contentTypeHeader {
	case "application/xml":
		if err = xml.Unmarshal(data, person); err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		if err = json.Unmarshal(data, person); err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	log.Printf("First name: %s\nLast name: %s\nPhone numbers: %v\n",
		person.FirstName, person.LastName, person.PhoneNumbers,
	)
	wr.WriteHeader(http.StatusOK)
}
