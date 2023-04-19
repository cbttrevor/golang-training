package main

import (
	"fmt"
	"os"
	"strconv"
)

func addNums(myvalue01 int, myvalue02 int) {
	fmt.Println("Total: ", (myvalue01 + myvalue02))
}

func main() {
	fmt.Println("Hello from CBT Nuggets")
	fmt.Println(2 + 2)
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	val1, _ := strconv.Atoi(os.Args[1])
	val2, _ := strconv.Atoi(os.Args[2])

	var myval1 int
	var myval2 int
	ret, _ := fmt.Scanln(&myval1)
	fmt.Println("Return value is: ", ret)
	ret2, _ := fmt.Scanln(&myval2)
	fmt.Println("Return value is: ", ret2)
	fmt.Println("Total is:", (myval1 + myval2))

	addNums(val1, val2)
}
