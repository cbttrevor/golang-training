package main

import (
	"fmt"

	"github.com/aerith86/trevmod"
	"github.com/aerith86/trevmod/github"
	"github.com/aerith86/trevmod/people"
)

func main() {
	trevmod.GetSomeData()
	github.DoGitHubStuff()
	john := people.Person{FirstName: "John", LastName: "McGovern"}
	fmt.Println(john)
	fmt.Println(people.Trevor)
	john.Walk(35)
	fmt.Printf("%+v\n", john)
}
