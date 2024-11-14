// Finding area using constant values

package main

import "fmt"

// Declare constants for Length and Width
const (
	LENGTH = 10
	WIDTH  = 5
)

func main() {
	// Declare area variable and calculate area
	var area int
	area = LENGTH * WIDTH

	// Print the area
	fmt.Printf("Value of area: %d\n", area)
}
