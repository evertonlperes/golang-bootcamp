package main

import "fmt"

// interface implementation of type string
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Custom logic for generating an english greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	// Custom logic for generating an spanish greeting
	return "Hola Companheiro!"
}
