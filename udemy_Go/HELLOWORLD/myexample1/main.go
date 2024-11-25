package main

import (
	"myexample1/mypackage"
)

func main() {
	// Call the Greet function from mypackage
	mypackage.Greet()

	// Call the Add function from my package
	mypackage.Add(57, 43)
}
