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
	var integer32 int = 2147483647
	fmt.Println(integer32)
}
