// Finding the length of an array

package main

import "fmt"

func main() {
	arr1 := [4]string{"Vol1", "Vol2", "Vol3", "Vol4"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println("Length of arr1:", len(arr1))
	fmt.Println("Length of arr2:", len(arr2))
}
