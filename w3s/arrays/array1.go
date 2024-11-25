// Decleare two arrays (arr1 and arr2) with inferred values

package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
