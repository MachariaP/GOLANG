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

// Add performs addation and adds some quirky commentary
func Add(a, b int) {
	sum := a + b
	fmt.Printf("Alright! You just added %d and %d and the result is... %d!\n", a, b, sum)
	if sum%2 == 0 {
		fmt.Println("Wow, that's an even number! Even numbers rock!")
	} else {
		fmt.Println("Odd number alert! Odd numbers are cool too!")
	}
}

// GetMessage returns a fun message that also gives a little joke
func GetMessage() string {
	// List of jokes
	jokes := []string{
		"Why don't skeletons fight each other? They don't have the guts!",
		"What do you call fake spaghetti? An impasta!",
		"Why was the math book sad? Because it had too many problems!",
	}
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]
	return joke
}
