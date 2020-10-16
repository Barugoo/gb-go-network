package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Home         Address  `json:"address"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

type Address struct {
	Street   string `json:"streetAddress"`
	City     string `json:"city"`
	PostCode string `json:"postalCode"`
}

func main() {
	person := Person{
		FirstName: "Иван",
		Home: Address{
			Street:   "Московское ш., 101, кв.101",
			City:     "Ленинград",
			PostCode: "101101",
		},
	}

	obj, err := json.Marshal(person)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s", obj)
}
