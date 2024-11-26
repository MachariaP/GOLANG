package main

// main is the entry point of the program
func main() {
	// create a slice of strings named "cards"
	// initialize it with "Ace of Diamonds
	// and the result of newCard function
	cards := deck{"Ace of Diamonds", newCard()}

	// append "Six of Spade" to the slice "cards"
	cards = append(cards, "Six of Spades")

	cards.print()

}

// newCard returns a string representing a card.
func newCard() string  {
	return "Five of Diamonds"
}
