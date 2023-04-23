package main

import "fmt"

func main() {
	var doLater = func() {
		fmt.Println("Doing stuff later")
	}
	printSomeData(doLater)

	func01 := returnSomeFunc(1)
	func02 := returnSomeFunc(2)
	func03 := returnSomeFunc(4)

	func01()
	func02()
	func03()
}

func returnSomeFunc(whichFunc uint) func() {
	if whichFunc == 1 {
		return func() {
			fmt.Println("This is func01")
		}
	}
	if whichFunc == 2 {
		return func() {
			fmt.Println("This is func02")
		}
	}
	if whichFunc == 3 {
		return func() {
			fmt.Println("This is func03")
		}
	}
	return func() {
		fmt.Println("This is default func")
	}
}

func printSomeData(doThisLater func()) {
	fmt.Println("Here is some info")
	doThisLater()
}
