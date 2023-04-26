package main

import "fmt"

type Pets interface {
	Dog | Cat | Bird
}
type Person[X Pets] struct {
	firstName string
	lastName  string
	pet       X
}

type Dog struct {
	dogBreed string
}

type Cat struct {
	catBreed string
}

type Bird struct {
	birdType string
}

func main() {
	dog01 := Dog{dogBreed: "Border Collie"}
	person01 := Person[Dog]{firstName: "Trevor", lastName: "Sullivan"}
	person01.pet = dog01
	fmt.Printf("%+v\n", person01)

	cat01 := Cat{catBreed: "Siamese"}
	person02 := Person[Cat]{firstName: "Trevor", lastName: "Sullivan"}
	person02.pet = cat01
	fmt.Printf("%+v\n", person02)

	bird01 := Bird{birdType: "Meadowlark"}
	person03 := Person[Bird]{firstName: "Trevor", lastName: "Sullivan"}
	person03.pet = bird01
	fmt.Printf("%+v\n", person03)
}
