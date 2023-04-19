package main

import (
	"fmt"
	"math"
)

var bankBalance int64 = 32769

func Rune() {
	var myrune rune = 'a'
	fmt.Print(myrune)
}

func StringStuff() {
	firstName := "Trevor"
	lastName := "Sullivan"
	company := "CBT Nuggets"

	var myFirstName string = "John"
	fmt.Println(myFirstName)

	fmt.Println("First name is: " + firstName + ", last name is: " + lastName + ", company is: " + company)
}

func main() {
	StringStuff()
	// Rune()
	// fmt.Println(bankBalance)
	// DoStuff()
	// Temperature()
	// humidity := 60
	// Humidity(&humidity)
	// fmt.Print("Humidity from main(): ", humidity, "\n")
}

func Humidity(myhumidity *int) {
	*myhumidity = 75
	fmt.Print("Humidity: ", *myhumidity, "\n")
}

func DoStuff() {
	// bankBalance = 268
	// fmt.Println(bankBalance)
	var newbankBalance int16
	if bankBalance > 100 {
		fmt.Println("int16 version of balance is ", int16(bankBalance))
		newbankBalance = int16(bankBalance) - 50
	}
	fmt.Println(newbankBalance)
	fmt.Println(math.MaxInt16)
}

func Temperature() {
	var temp float64
	temp = 5.6
	tempDelta := 5.7
	temp = temp + tempDelta
	fmt.Println("Temperature is currently ", temp)

	fmt.Print("Max value of float32 is ", math.MaxFloat32, "\n")
	fmt.Print("Max value of float64 is ", math.MaxFloat64, "\n")
}
