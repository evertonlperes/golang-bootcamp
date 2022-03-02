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
	lastNAme  string
	contact   contactInfo
}

func main() {
	// # First way to use struct
	// wicks := person{firstName: "Wicks", lastNAme: "Testing"}
	// fmt.Println(wicks)

	// # Second way
	var wicks person

	wicks.firstName = "Wicks"
	wicks.lastNAme = "Testing"

	fmt.Println(wicks)       //output: { }
	fmt.Printf("%+v", wicks) // output: {firstName: lastNAme:}%
}
