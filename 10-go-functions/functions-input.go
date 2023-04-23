package main

import "fmt"

func main() {
	fmt.Println("Learning about Go functions")
	myPerson01 := Person{firstName: "John", lastName: "McGovern", age: 26}
	doSomeWork(myPerson01, false)
}

type Person struct {
	firstName string
	lastName  string
	age       uint
}

func doSomeWork(p Person, hasBrownHair bool) {
	fmt.Println("Calling a function")
	fmt.Println("The name you specified was " + p.firstName)
	fmt.Printf("This person's age is %v\n", p.age)
	fmt.Printf("Does %v have brown hair? %v", p.firstName, hasBrownHair)
}
