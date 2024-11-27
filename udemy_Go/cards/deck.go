package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings.
// A deck is essentially a collection of cards, where each card is represented as a string.
type deck []string

// newDeck creates a new deck of cards.
// It initializes a deck with 5 card values and 4 suits (Spades, Diamonds, Hearts, and Clubs).
// Each card in the deck is a combination of a value and a suit (e.g., "Ace of Spades").
func newDeck() deck {
	cards := deck{} // Initialize an empty deck (a slice of strings)

	// Define the suits and values for the cards
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	// Loop through each suit and value to create the cards and add them to the deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // Append the card to the deck
		}
	}

	return cards // Return the newly created deck
}

// print is a method on the 'deck' type that prints each card in the deck along with its index.
// It loops over the deck and prints both the index (position in the deck) and the card itself.
func (d deck) print() {
	for i, card := range d {
		// Print the index and the card value (e.g., "0 Ace of Spades")
		fmt.Println(i, card)
	}
}

// deal function takes a deck and a handSize as input.
// It returns two decks: the first deck represents the hand (the first 'handSize' cards),
// and the second deck represents the remaining cards after the hand is dealt.
func deal(d deck, handSize int) (deck, deck) {
	// Slice the original deck into two parts: hand and remaining cards
	return d[:handSize], d[handSize:]
}

// toString is a method on the 'deck' type that converts the deck into a single string.
// It joins the cards in the deck using commas and returns the resulting string.
func (d deck) toString() string {
	// Convert the deck to a string by joining the cards with a comma
	return strings.Join([]string(d), ",")
}
