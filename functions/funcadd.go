// add takes two arguments of type int

package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(43, 13))
}
