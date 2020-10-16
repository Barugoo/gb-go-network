package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName    string   `json:"firstName" xml:"firstName"`
	LastName     string   `json:"lastName"`
	Home         Address  `json:"address"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

type Address struct {
	Street   string `json:"streetAddress"`
	City     string `json:"city"`
	PostCode string `json:"postalCode"`
}

var obj = []byte(`{
"firstName": "Иван",
"lastName": "Иванов",
"address": {
"streetAddress": "Московское ш., 101, кв.101",
"city": "Ленинград",
"postalCode": "101101"
},
"phoneNumbers": [
"812 123-1234",
"916 123-4567"
]
}`)

func main() {
	person := Person{}

	if err := json.Unmarshal(obj, &person); err != nil {
		log.Panic(err)
	}

	print(person.FirstName)
}
