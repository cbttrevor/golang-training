package main

import (
	"fmt"
)

type Car struct {
	manufacturer string
	model        string
	moonroof     bool
	horsepower   uint16
	year         uint16
	owner        *Person
}

type Person struct {
	firstName      string
	lastName       string
	company        string
	age            uint8
	distanceWalked uint32
}

func (p *Person) Walk(distance uint8) {
	p.distanceWalked += uint32(distance)
}

func testDistanceWalked() {
	person01 := Person{firstName: "Billy", lastName: "Bob", age: 22, company: "Contoso"}

	person01.Walk(5)
	person01.Walk(12)
	fmt.Printf("%+v\n", person01)
}

func main() {
	person01 := Person{firstName: "Trevor", lastName: "Sullivan", company: "CBT Nuggets", age: 25}
	person02 := Person{firstName: "John", lastName: "McGovern", company: "CBT Nuggets", age: 27}

	person01.firstName = "John"
	person01.age = 32
	fmt.Printf("%v %v is my name, and I work for %v\n", person01.firstName, person01.lastName, person01.company)
	fmt.Printf("%v's age is %v\n", person01.firstName, person01.age)

	fmt.Printf("%+v\n", person01)
	fmt.Printf("%p\n", &person01)
	fmt.Printf("%p\n", &person02)

	// testStructRefs()
	testDistanceWalked()
}

func testStructRefs() {
	car01 := Car{manufacturer: "Hyundai", model: "Santa Fe", year: 2015, moonroof: true, horsepower: 220}
	carOwner := Person{firstName: "Daniel", lastName: "Williams", age: 42, company: "CBT Nuggets"}
	car01.owner = &carOwner

	carOwner.firstName = "Trevor"

	fmt.Printf("%+v\n", car01.owner)
	fmt.Printf("%p %p\n", &car01.owner, &carOwner)
}
