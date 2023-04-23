package main

import "fmt"

func main() {
	printSomeNames(26, "Trevor", "Nathan", "John")
	printSomeNames(27, "Trevor", "Nathan")
	printSomeNames(28, "Trevor")
}

func printSomeNames(age uint, names ...string) {
	for _, val := range names {
		fmt.Printf("Name is %v\n", val)
	}
}
