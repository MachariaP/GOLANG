// Expicit type conversion

package main

import "fmt"

func main() {
	var integer16 int16 = 127
	var integer32 int32 = 32767
	fmt.Println(int32(integer16) + integer32)
}
