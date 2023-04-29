package main

import (
	j "encoding/json"
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	Company     string
	Age         uint8
	Temperature float32
	Friends     []string
}

func main() {
	// example01()
	exampleMaps()
}

func exampleMaps() {
	var jsonstring string
	jsonstring = `
	{
		"FirstName": "Trevor",
		"LastName": "Sullivan",
		"Company": "CBT Nuggets",
		"Age": 32,
		"Temperature": 98.6,
		"Friends": [
			"Billy",
			"Rachel",
			"Joe"
		]
	}
	`

	if !j.Valid([]byte(jsonstring)) {
		fmt.Println("There was an error parsing the JSON document")
		return
	}

	var jsonresult map[string]int
	var employeeresult Employee = Employee{}

	err := j.Unmarshal([]byte(jsonstring), &employeeresult)

	fmt.Println(err)
	fmt.Println(jsonresult)
	fmt.Printf("%+v\n", employeeresult)

	for k, v := range jsonresult {
		fmt.Printf("Key: %v Value: %v\n", k, v)
	}

	fmtString := "%v %v is age %v and his friends are: %v"
	fmt.Printf(fmtString, employeeresult.FirstName, employeeresult.LastName, employeeresult.Age, employeeresult.Friends)
}

func example01() {
	var jsonstring string
	jsonstring = `
	["Trevor Sullivan is a CBT Nuggets trainer."]
	`

	jsonstring = `
	["Trevor", "John", "Nathan", "Daniel", "Trevor", "John", "Nathan", "Daniel"]
	`

	jsonstring = `
	[37, 236, 9992, true, false, "Trevor", "John", "Nathan", "Daniel", "Trevor", "John", "Nathan", "Daniel"]
	`

	if !j.Valid([]byte(jsonstring)) {
		fmt.Println("Your JSON wasn't valid. Aborting operation.")
		return
	}

	var jsonresult []any

	err := j.Unmarshal([]byte(jsonstring), &jsonresult)
	if err != nil {
		fmt.Println("There was a problem decoding your JSON document. Please inspect the source data and try again.")
		fmt.Println(err)
	} else {
		fmt.Println("Your JSON decoding was successful.")
		fmt.Println(jsonresult)
	}

	for _, value := range jsonresult {
		fmt.Printf("Value is %v\n", value)
	}
}
