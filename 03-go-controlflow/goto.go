package main

import "fmt"

func testGoto() {
	var input int
StartPrint:
	fmt.Println("Welcome to CBT Nuggets!")
	fmt.Println("Welcome to Golang training!")
	fmt.Scanln(&input)
	if input == 1 {
		goto StartPrint
	}
}
