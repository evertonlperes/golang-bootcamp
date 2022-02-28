package main

func main() {

	cards := newDeck()

	// hand = deck
	// remainingCard = deck (remaning cards on the deck)
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
