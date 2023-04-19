package main

import "fmt"

func testSwitchInt() {
	var bankBalance int

	fmt.Println("What is your balance?")
	fmt.Scanln(&bankBalance)

	switch {
	case bankBalance > 50000:
		fmt.Println("You can purchase an exotic car!")
	case bankBalance > 10000:
		fmt.Println("You can purchase a regular car!")
	default:
		fmt.Println("You can't purchase any car!")
	}
}

func testSwitch() {
	var carBrand string
	fmt.Println("What car brand do you want?")
	fmt.Scanln(&carBrand)

	switch carBrand {
	case "Hyundai":
		fmt.Println("You can purchase a Hyundai!")
	case "Toyota":
		fmt.Println("You can purchase a Toyota!")
	default:
		fmt.Println("You can't purchase that brand!")
	}
}

func testIf() {
	var carBrand string

	fmt.Println("What car brand do you want?")
	fmt.Scanln(&carBrand)

	if carBrand == "Hyundai" || carBrand == "Honda" || carBrand == "Tesla" {
		fmt.Println("You can purchase a car!")
	} else {
		fmt.Println("You can't buy a car!")
	}
}
