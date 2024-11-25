// Using the strconv package for convertion

package main

import (
	"fmt"
	"strconv" // the strconv package for conversion
)

func main() {
	// convert a string to an integer using strconv.Atoi
	// strconv.Atoi returns an integer from a string
	// the second return value is an error value is returned
	// if the string is not a valid integer value then an error is returned

	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return
	}

	// Convert an integer to a string using strconv.Itoa
	// strconv.Itoa returns a string from an integer

	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
