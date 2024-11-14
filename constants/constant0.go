// Declaring a Constant in Go
package main

import (
	"fmt"
)

// Declare a constant named 'PI' with a value of 3.14
const PI = 3.14

func main() {
	// Print the value of the constant 'PI'
	fmt.Println(PI)

	// Print the type of the constant 'PI'
	fmt.Printf("Type %T\n", PI)

	// Print the value of the constant 'PI' with a precision of 2
	fmt.Printf("Value %.2f\n", PI)
}
