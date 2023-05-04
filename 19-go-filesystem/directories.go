package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var folderName string = "gotestdata"
	// err := os.Mkdir(folderName, 0755)

	// var folderStructure string = "parent/child01/child02"
	// err := os.MkdirAll(folderStructure, 0755)
	// if err == nil {
	// 	fmt.Println(err)
	// }

	// err := os.Remove("gotestdata")
	// if err != nil {
	// 	if err == err.(*os.PathError) {
	// 		fmt.Println("Type assertion succesful. This is a path error.")
	// 	}
	// 	if strings.Contains(err.Error(), "no such file") {
	// 		fmt.Printf("There wasn't a folder or file found with the name %v", "gotestdata")
	// 	}
	// 	fmt.Println(err)
	// }

	// CreateFile()
	// ReadTextLines()

	DoFileAdvancedStuff()
}

func DoFileAdvancedStuff() {
	fileName := "cars.txt"
	carfile, _ := os.Open(fileName)
	defer carfile.Close()

	seekDistance, _ := carfile.Seek(6, 0)
	fmt.Printf("I seeked %v bytes\n", seekDistance)

	buffer := make([]byte, 6)
	readBytes, _ := carfile.Read(buffer)
	fmt.Printf("I read %v bytes from file\n", readBytes)
	fmt.Printf("The data read was: %v\n", string(buffer))

	buffer = make([]byte, 4)
	carfile.Seek(7, 1)
	readBytes, _ = carfile.Read(buffer)

	fmt.Printf("I read %v bytes from file\n", readBytes)
	fmt.Printf("The data read was: %v\n", string(buffer))
}

func ReadTextLines() {
	fileName := "cars.txt"
	carfile, _ := os.Open(fileName)
	textScanner := bufio.NewScanner(carfile)
	for textScanner.Scan() {
		fmt.Printf("My favorite car brand is %v\n", textScanner.Text())
	}
}

func CreateFile() {
	fileName := "cars.txt"
	fileContent := `Honda
Toyota
Buick
Jeep`
	err := os.WriteFile(fileName, []byte(fileContent), 0750)
	fmt.Println(err)
}
