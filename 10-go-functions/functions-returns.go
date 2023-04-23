package main

import "fmt"

func main() {
	fmt.Printf("The age that was returned is %v\n", getAge("Trevor"))
	fmt.Printf("The age that was returned is %v\n", getAge("John"))
	fmt.Printf("The age that was returned is %v\n", getAge("Daniel"))

	age01, hair01 := getDetails("Trevor")
	fmt.Printf("The person's age is %v and hair is color %v", age01, hair01)
}

func getAge(firstName string) uint {
	if firstName == "Trevor" {
		return 35
	} else if firstName == "John" {
		return 34
	}
	return 0
}

func getDetails(firstName string) (age uint, hairColor string) {
	if firstName == "Trevor" {
		age = 35
		hairColor = "Blue"
		return
	} else if firstName == "John" {
		age = 34
		hairColor = "Brown"
		return
	}
	return 0, "None"
}
