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
	contactInfo
}

func main() {

	wicks := person{
		firstName: "Wicks",
		lastName:  "Testing",
		contactInfo: contactInfo{
			email:   "test@gmail.com",
			zipCode: "9000",
		},
	}

	wicksPointer := &wicks // &: give the memory address fo the value this variable is pointing at
	wicksPointer.updateName("Aloww")
	wicks.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // *pointer: give the value this memory address is pointing at
}

// Function to print the struct content
func (p person) print() {
	fmt.Printf("%+v", p)
}
