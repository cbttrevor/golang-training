package main

import "fmt"

func testFor() {
	var carBrand string = "Honda"
	for carBrand == "Honda" {
		fmt.Println("For executed")
		carBrand = "Toyota"
	}
}

func testForRange() {
	var firstName string = "John"

	for i, value := range firstName {
		fmt.Println(value)
		if value == 'J' {
			fmt.Println("Found a capital J")
		}
		fmt.Println("The index of this value is ", i)
	}
}

func testForRangeArray() {
	var bankBalances [5]int
	bankBalances[0] = 500
	bankBalances[1] = 50500
	bankBalances[2] = 75000
	bankBalances[3] = 150000
	bankBalances[4] = 200000

	for _, balance := range bankBalances {
		fmt.Println(balance)
		if balance > 60000 {
			fmt.Println("This dude is rich!")
		} else {
			fmt.Println("You need to keep working")
		}
	}

}

func testForRangeMap() {
	var foods = make(map[int]string)
	foods[0] = "bacon"
	foods[1] = "eggs"
	foods[2] = "french toast"

	for _, food := range foods {
		fmt.Println(food)
		if food == "french toast" {
			fmt.Println("That is my favorite!")
		}
	}

}
