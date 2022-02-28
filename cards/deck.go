package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Helper: Convert deck to string
func (d deck) toString() string {

	// convert d to type string slice
	// and separate it by comma
	return strings.Join([]string(d), ",")

}

// Save cards into a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read the cards from saved file
// and return a new Deck
func newDeckFromFile(filename string) deck {
	//bs = byteSlice
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error --> ", err)
		os.Exit(1)
	}

	// convert byteSlice to string and attribute it to 's'
	s := strings.Split(string(bs), ",")
	// convert 's' to a deck
	return deck(s)
}

// Shuffle the cards
func (d deck) shuffle() {

	// creating a rand seed based on the current
	// execution time
	source := rand.NewSource(time.Now().UnixNano())
	// 'r' = random
	r := rand.New(source)

	for i := range d {
		// 'np' = new position
		// random set number based on the size of the string slice
		np := r.Intn(len(d) - 1)
		// move elements into the array
		d[i], d[np] = d[np], d[i]
	}
}
