package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string   `xml:"firstName" json:"firstName"`
	LastName     string   `xml:"lastName"`
	Home         Address  `xml:"home"`
	PhoneNumbers []string `xml:"phones"`
}

type Address struct {
	Street   string `xml:"street"`
	City     string `xml:"city"`
	PostCode string `xml:"postCode"`
}

func main() {
	person := Person{
		FirstName: "Иван",
		LastName:  "Иванов",
		Home: Address{
			Street:   "Московское ш., 101, кв.101",
			City:     "Ленинград",
			PostCode: "101101",
		},
		PhoneNumbers: []string{
			"812 123-1234",
			"916 123-4567",
		},
	}

	obj, err := xml.Marshal(person)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("%s", obj)
}
