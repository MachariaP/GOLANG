// circleArea calculates the area of a circle given its radius.
// It takes a float64 value (radius) as input and returns the area as a float64.
// Formula used: Area = π * radius^2

package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Printf("Good morning! %v\n", n)
}

func sayGoodbye(n string) {
	fmt.Printf("Goodbye! %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	names := []string{"John", "Anne", "Jack"}
	cycleNames(names, sayGreetings)
	cycleNames(names, sayGoodbye)
	a1 := circleArea(10.5)
	a2 := circleArea(15.0)

	fmt.Println(a1, a2)

	//cycleNames([]string{"John", "Anne", "Jack"}, sayGreetings)
	//cycleNames([]string{"John", "Anne", "Jack"}, sayGoodbye)
}
