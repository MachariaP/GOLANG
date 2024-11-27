package main

import "fmt"

// main is the entry point of the program.
// In this function, a new deck of cards is created using the newDeck function,
// and then the deck is converted to a string and printed to the console.
func main() {
	// Create a new deck using the newDeck function.
	cards := newDeck()

	// Convert the deck to a string and print it to the console.
	// The cards.toString() method will join the cards in the deck with commas.
	fmt.Println(cards.toString())
}
