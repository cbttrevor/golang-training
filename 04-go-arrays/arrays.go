package main

import (
	"fmt"
	"reflect"
	"sort"
)

func testIntArray() {
	var raceTimes [4]int

	for i := range raceTimes {
		fmt.Print("Enter a new race time: ")
		fmt.Scanln(&raceTimes[i])
	}

	for _, raceTime := range raceTimes {
		fmt.Print("Race time is ")
		fmt.Print(raceTime, " seconds\n")
	}

	fmt.Println("Array length: ", len(raceTimes))
	fmt.Println("Array values: ", raceTimes)
	fmt.Println("Type of raceTimes is: ", reflect.TypeOf(raceTimes).Kind())
}

func main() {
	// testIntArray()
	// testStringArray()
	// testGoSlice()
	testGoSliceString()
}

func testGoSliceString() {
	var nameList []string
	var newName string

	for {
		fmt.Print("Please enter a new name: ")
		fmt.Scanln(&newName)

		if newName == "done" {
			break
		}
		nameList = append(nameList, newName)
	}

	sort.Strings(nameList)

	for _, name := range nameList {
		fmt.Println(name)
	}
	fmt.Println("You entered", len(nameList), "names")
}

func testStringArray() {
	var contestants = [5][3]string{{"Trevor", "Sullivan", "39"}, {"John", "McGovern", "36"}}

	contestants[2] = [3]string{"Keith", "Barker", "35"}
	fmt.Println(contestants)
}

func testGoSlice() {
	var raceTimes [12]int

	raceTimes[0] = 5
	raceTimes[5] = 583
	raceTimes[11] = 892

	mySliceTimes := raceTimes[0:6]
	fmt.Println(mySliceTimes)

	fmt.Println(raceTimes)
	fmt.Println("Length of slice is: ", len(raceTimes))
	fmt.Println("The type of raceTimes is: ", reflect.TypeOf(raceTimes).Kind())
	fmt.Println("The type of mySliceTimes is: ", reflect.TypeOf(mySliceTimes).Kind())
	if reflect.TypeOf(raceTimes).Kind() == reflect.Slice {
		fmt.Println("We have confirmed that this raceTimes is a slice, not an array")
	}
}
