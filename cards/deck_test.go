package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// 'd' new deck
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	firstItem := d[0]
	if firstItem != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", firstItem)
	}

	lastItem := d[len(d)-1]
	if lastItem != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", lastItem)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	testFile := "_decktesting"
	// make sure the test file is gone
	os.Remove(testFile)

	deck := newDeck()
	deck.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	// final teardown after tests
	os.Remove(testFile)
}
