// Static type declaration in Go
/* It enables the compiler to check the type correctness of the program at compile time. */

package main

import (
	"fmt"
)

func main() {
	// Declare a variable 'x' of type float64
	var x float64

	// Assign a value to the variable 'x'
	x = 20.0

	// Print the value of the variable 'x'
	fmt.Println(x)

	// Print the type 'x' using the Printf function with the %T format specifier
	fmt.Printf("x is of type %T\n", x)
}
