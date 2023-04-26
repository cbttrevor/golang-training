package main

import "fmt"

type FirstName string
type DogBreed string

func Add[A uint16 | uint32 | uint64, B string | DogBreed](x, y A, inputString B) {
	fmt.Printf("%v %v %v\n", x, y, inputString)
}

func main() {

	// var inputS string = "CBT Nuggets"
	// var dbreed DogBreed = "Border Collie"
	var fName FirstName = "Trevor"

	var x, y uint16
	x = 56
	y = 89
	Add(x, y, fName)

	var x2, y2 uint64
	x2 = 5372
	y2 = 5326
	Add(x2, y2, fName)
}
