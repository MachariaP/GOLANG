// Changing the value of the third element in the prices array

package main

import "fmt"

func main() {
	prices := [3]int{10, 20, 30}

	prices[2] = 55

	fmt.Println(prices)
}
