package main

import "fmt"

// Tip: https://play.golang.org/

func main() {
	// types in go is strict, must define the type.
	// var card string = "Ace of Spades"
	card := newCard() // Another way to define a variable in go (recommended)
	// card = "Five of Diamonds" // Replacing the first value to Five of Diamonds

	fmt.Println(card)
}

// Go requires to give a type for the function, in this case it will return a type string
func newCard() string {
	return "Five of Diamonds"
}
