// Set a variable of type int and use the value from the integer32 or integer64 variable
// to confirm the natural size of the variable on your system.
// If you are on a 32-bit system and use a value higher than the default (2,147,483,647),
// you will get an overflow error tha looks like this:
// runtime error: constant 9223372036854775807 overflows int

package main

import (
	"fmt"
)

func main() {
	// 32-bit integer test
	var integer32 int = 2147483647
	fmt.Println("32-bit integer: ", integer32)

	// 64-bit integer test
	var integer64 int = 9223372036854775807
	fmt.Println("64-bit integer: ", integer64)

	// Test with a value that exceeds the 32-bit integer limit
	var integerOverflow int64 = 2147483648
	fmt.Println("64-bit integer (overflow test): ", integerOverflow)
}
