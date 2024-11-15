// The bool type represents a boolean and its either true or false

package main

import (
	"fmt"
)

func main() {
	a, b := true, false
	c := a && b
	d := a || b

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
