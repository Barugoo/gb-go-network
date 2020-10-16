package main

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
