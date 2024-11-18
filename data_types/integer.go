// Various integers in Go

package main

import "fmt"

func main() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Printf("%d %d %d %d\n", integer8, integer16, integer32, integer64)
}
