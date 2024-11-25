package main

import (
	"fmt"
	"myexample1/mypackage"
)

func main() {
	// Call the Greet function from mypackage
	mypackage.Greet()

	// Call the Add function from my package
	mypackage.Add(57, 43)

	// Get a fun message (joke) from mypackage
	joke := mypackage.GetMessage()
	fmt.Println("Here's a joke for you: " + joke)
}
