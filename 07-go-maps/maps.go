package main

import "fmt"

func testGoMapSimple() {
	mySimpleMap := map[uint8]uint8{}
	mySimpleMap[167] = 255
	mySimpleMap[132] = 145
	mySimpleMap[45] = 12
	fmt.Printf("%v\n", mySimpleMap)
	fmt.Printf("%v\n", len(mySimpleMap))

	// NOTE: Prompt a user for a specific Map key to retrieve
	fmt.Printf("Please enter a key to retrieve: ")
	var keyId uint8
	fmt.Scanln(&keyId)
	_, keyExists := mySimpleMap[keyId]
	if keyExists {
		fmt.Printf("The user entered: %v\n", mySimpleMap[keyId])
	} else {
		fmt.Printf("That key wasn't found in the map. Try again.\n")
	}

	for myKey, myValue := range mySimpleMap {
		if myKey > 150 {
			fmt.Println("Key is: ", myKey)
			fmt.Println("Value is: ", myValue)
		}
	}
}

func testGoMapStringKey() {
	myNameMap := make(map[string]string, 20)

	var myKey, myValue string
	for {
		fmt.Printf("Please enter a key: ")
		fmt.Scanln(&myKey)
		if myKey == "done" {
			break
		}
		fmt.Printf("Please enter a value: ")
		fmt.Scanln(&myValue)
		myNameMap[myKey] = myValue

	}

	fmt.Printf("Length of myNameMap is: %v\n", len(myNameMap))
	fmt.Printf("The value for your key %v was %v", myKey, myValue)

}

func testGoMapArrayKey() {
	myArray := [2]string{"Trevor", "Sullivan"}
	myVehicleMap := make(map[[2]string][5]string)
	myVehicleMap[myArray] = [5]string{"Honda", "Tesla", "Lambo"}

	john := [2]string{"John", "McGovern"}
	myVehicleMap[john] = [5]string{"Hyundai", "Ford", "Opal", "Honda", "Toyota"}
	fmt.Printf("%v\n", myVehicleMap)

	mySliceMap := make(map[[2]string][]string)
	mySliceMap[[2]string{"Trevor", "Sullivan"}] = []string{"Honda"}
	fmt.Printf("%v\n", mySliceMap)
}

func main() {
	fmt.Println("Learning about Go Maps")

	fmt.Println("Maps are key-value pairs.")
	fmt.Println("You can use any Go data type for the key (except Slice) and the value")
	fmt.Println("Go Maps also known as HashTables or Dictionaries in other programming languages.")

	// testGoMapSimple()
	// testGoMapStringKey()
	testGoMapArrayKey()
}
