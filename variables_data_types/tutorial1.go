// Dynamic Type Declaration / Interface Declaration
/* Requires the compiler to interplate the type of the variable based
on the value passed to it */

package main

import (
	"fmt"
)

func main() {
	// Declare a variable 'x' of type float64 and assign it to 20.0
	var x float64 = 20.0

	// Declare a variable 'y' and assign it to 42
	y := 42

	// Print the value of the variable 'x' and 'y'
	fmt.Println(x)
	fmt.Println(y)

	// Print the type 'x' and 'y' using the Printf function with the %T format specifier
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
}
