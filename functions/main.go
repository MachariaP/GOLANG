package main

import "fmt"

func sayGreetings(n string) {
	fmt.Printf("Good morning! %v\n", n)
}

func getResults(n string) {
	fmt.Printf("Where is your result slip? %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye! %v\n", n)
}

func main() {
	sayGreetings("John")
	getResults("John")
	sayBye("John")
}
