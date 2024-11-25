package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Strict typing: Declare variables with specific types
	var integer int = 10
	var floatNum float64 = 20.5
	var boolean bool = true
	var text string = "Hello, Go!"

	// Default values: Variables not initialized will have default values
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	// Displaying values and default values
	fmt.Println("Integer:", integer)
	fmt.Println("Float:", floatNum)
	fmt.Println("Boolean:", boolean)
	fmt.Println("String:", text)

	fmt.Println("Default Integer:", defaultInt)
	fmt.Println("Default Float:", defaultFloat)
	fmt.Println("Default Boolean:", defaultBool)
	fmt.Println("Default String:", defaultString)

	// Type conversion: Convert one type to another
	var smallInt int16 = 100
	var largeInt int32 = 200
	sum := int32(smallInt) + largeInt // Manual conversion
	fmt.Println("Sum after type conversion:", sum)

	// Using strconv package for type conversion
	numStr := "42"
	num, err := strconv.Atoi(numStr) // Convert string to int
	if err == nil {
		fmt.Println("Converted string to int:", num)
	}

	newStr := strconv.Itoa(num) // Convert int back to string
	fmt.Println("Converted int back to string:", newStr)
}
