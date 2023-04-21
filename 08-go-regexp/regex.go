package main

import (
	"fmt"
	"regexp"
)

func testRegexMatch() {
	var regexPattern = "^Trevor"
	regexPattern = "Kubernetes\\$$"
	var inputString = "Trevor Sullivan is learning about PowerShell. Billy is learning about Kubernetes$"
	didItMatch, _ := regexp.MatchString(regexPattern, inputString)
	// fmt.Println("Did the pattern match the input string? ", didItMatch)
	if didItMatch {
		fmt.Println("Your string ended with Kubernetes$.")
	} else {
		fmt.Println("Your string did NOT end with Kubernetes$.")
	}
}

func main() {
	// testRegexMatch()
	// testRegexQuantifiers()
	testNumberOfGroceries()
}

func testNumberOfGroceries() {
	var numGroceriesPattern = "[[:digit:]]{2,}"
	var namePattern = "^[[:alpha:]]{2,}"
	var inputString = "John Sullivan went to the store and bought 878 groceries."
	regex, _ := regexp.Compile(numGroceriesPattern)
	nameRegex, _ := regexp.Compile(namePattern)
	numGroceries := regex.FindString(inputString)
	firstName := nameRegex.FindString(inputString)
	fmt.Printf("The number of groceries purchased was %v\n", numGroceries)
	fmt.Printf("The person who retrieved groceries was %v\n", firstName)
}

func testRegexQuantifiers() {
	var quantifierPattern = "[[:digit:]]{2,}"
	var inputString = "Trevor Sullivan went to the store for groceries."
	inputString = "Trevor Sullivan went to the store and bought 2 groceries."
	didItMatch, _ := regexp.MatchString(quantifierPattern, inputString)
	fmt.Printf("%v\n", didItMatch)
}
