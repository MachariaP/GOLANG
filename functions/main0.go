/*
This script demonstrates the use of higher-order functions in Go.
It includes three functions:

1. `sayGreetings`: Takes a name (string) as input and prints a greeting message.
2. `sayGoodbye`: Takes a name (string) as input and prints a goodbye message.
3. `cycleNames`: Accepts a slice of names (strings) and a function (f) that takes a string argument.
   It iterates over the names and applies the passed function to each name.

The `main` function creates a list of names and demonstrates how the `cycleNames` function can be used
with both the `sayGreetings` and `sayGoodbye` functions to print messages for each name.

*/

package main

import "fmt"

func sayGreetings(n string) {
	fmt.Printf("Good morning! %v\n", n)
}

func sayGoodbye(n string) {
	fmt.Printf("Goodbye! %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	names := []string{"John", "Anne", "Jack"}
	cycleNames(names, sayGreetings)
	cycleNames(names, sayGoodbye)

	//cycleNames([]string{"John", "Anne", "Jack"}, sayGreetings)
	//cycleNames([]string{"John", "Anne", "Jack"}, sayGoodbye)
}
