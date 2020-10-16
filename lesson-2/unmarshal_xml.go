package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string   `xml:"firstName"`
	LastName     string   `xml:"lastName"`
	Home         Address  `xml:"address"`
	PhoneNumbers []string `xml:"phoneNumbers"`
}

type Address struct {
	Street   string `xml:"streetAddress"`
	City     string `xml:"city"`
	PostCode string `xml:"postalCode"`
}

var obj = `<Person><FirstName>Иван</FirstName><LastName>Иванов</LastName><Home><Street>Московское ш., 101, кв.101</Street><City>Ленинград</City><PostCode>101101</PostCode></Home><PhoneNumbers>812 123-1234</PhoneNumbers><PhoneNumbers>916 123-4567</PhoneNumbers></Person>`

func main() {
	person := Person{}

	if err := xml.Unmarshal([]byte(obj), &person); err != nil {
		log.Panicln(err)
	}

	fmt.Println(person.FirstName)
}
