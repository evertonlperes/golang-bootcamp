package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	// Splitting it on suites and values
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Iterate over suites and values
	// and appends it to a card
	for _, suit := range cardSuites { // use '_' to tell go to ignore the variable not used.
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print the cards from the slicer
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Make a deal and split cards based on the hand size
// also it returns two values (left:right from slicers)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Helper deck to string, before convert it to byte slice
func (d deck) toString() string {

	// convert d to type string slice
	// and separate it by comma
	return strings.Join([]string(d), ",")

}
