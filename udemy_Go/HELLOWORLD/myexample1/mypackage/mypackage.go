// mypackage.go

package mypackage

import (
	"fmt"
	"math/rand"
	"time"
)

// Greet generates and prints a random greeting message from a predefined list.
func Greet() {
	// List of possible greeting messages
	greetings := []string{
		"Hello there! Ready for some fun?",
		"Hey! Let's make this interesting!",
		"Hi, friend! Let's do something awesome!",
	}

	// Initialize the random number generator to ensure different results
	rand.Seed(time.Now().UnixNano())

	// Select a random greeting from the list
	greeting := greetings[rand.Intn(len(greetings))]

	// Print the selected greeting
	fmt.Println(greeting)
}
