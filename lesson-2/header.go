package main

import (
	"encoding/json"
	"encoding/xml"
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

	log.Fatal(http.ListenAndServe(":8091", router))
}

func personHandler(wr http.ResponseWriter, req *http.Request) {
	acceptHeader := req.Header.Get("Accept")

	person := &Person{
		FirstName: "Dmitry",
		LastName:  "Shelamov",
		PhoneNumbers: []string{
			"88005553535",
		},
	}
	var respBody []byte
	var err error
	switch acceptHeader {
	case "application/xml":
		respBody, err = xml.Marshal(person)
		if err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		wr.Header().Set("Content-Type", "application/xml")
	default:
		respBody, err = json.Marshal(person)
		if err != nil {
			log.Println(err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		wr.Header().Set("Content-Type", "application/json")
	}
	wr.Write(respBody)
}
