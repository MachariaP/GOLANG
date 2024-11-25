// Creating a slice from an array

package main

import "fmt"

func main() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Println("myslice = %v\n", myslice)
	fmt.Println("length = %d\n", len(myslice))
	fmt.Println("capacity = d\n", cap(myslice))
}
