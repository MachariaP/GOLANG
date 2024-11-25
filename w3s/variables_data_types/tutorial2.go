// Mixed Variable Declaration
/* Declare multiple variables of different types in a single line */

package main

import (
	"fmt"
)

func main() {
	// Declare multiple variables of different types in a single line
	var a, b, c = 3, 4, "foo"

	// Print the value of the variables 'a', 'b' and 'c'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Print the type of the variables 'a', 'b' and 'c' using the Printf function with the %T format specifier
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}
