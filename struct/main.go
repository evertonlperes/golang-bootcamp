package main

import "fmt"

// Define the struct: contact Info
type contactInfo struct {
	email   string
	zipCode string
}

// Define the struct: person
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	wicks := person{
		firstName: "Wicks",
		lastName:  "Testing",
		contact: contactInfo{
			email:   "test@gmail.com",
			zipCode: "9000",
		},
	}

	fmt.Printf("%+v", wicks)
}
