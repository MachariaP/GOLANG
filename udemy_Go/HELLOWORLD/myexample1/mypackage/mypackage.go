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

func Add(a, b int) {
	sum := a + b
	fmt.Printf("Alright! You just added %d and %d and the result is... %d!\n", a, b, sum)
	if sum%2 == 0 {
		fmt.Println("Wow, that's an even number! Even numbers rock!")
	} else {
		fmt.Println("Odd number alert! Odd numbers are cool too!")
	}
}
