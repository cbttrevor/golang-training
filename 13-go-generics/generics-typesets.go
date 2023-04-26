package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type FirstName string
type DogBreed string

func Add[A constraints.Unsigned, B constraints.Ordered](x, y A, inputString B) {
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

	// var x3, y3 float32
	// x3 = 5372.56
	// y3 = 5326.2
	// Add(x3, y3, dbreed)

	// var x4, y4 float64
	// x4 = 5372.56
	// y4 = 5326.2
	// Add(x4, y4, dbreed)
}
