package main

import "fmt"

// main is the entry point of the program
func main() {
	// create a slice of strings named "cards"
	// initialize it with "Ace of Diamonds
	// and the result of newCard function
	cards := []string{"Ace of Diamonds", newCard()}

	// append "Six of Spade" to the slice "cards"
	cards = append(cards, "Six of Spades")

	// Loop over the slice "cards", using range to get both index(i) 
	// and value (card)
	for i, card := range cards {

		// print the index and the card value
		fmt.Println(i, card)
	}
}

// newCard returns a string representing a card.
func newCard() string  {
	return "Five of Diamonds"
}
