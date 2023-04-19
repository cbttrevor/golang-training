package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func testFormatSpecifiers() {
	firstName := "Trevor"
	lastName := "Sullivan"
	course := "Kubernetes"
	fmt.Printf("%v %v is learning about %x", firstName, lastName, course)

	myHexString := "4b756265726e65746573"
	decodedBytes, _ := hex.DecodeString(myHexString)
	decodedString := string(decodedBytes)
	fmt.Printf("%v\n", decodedString)
}

func testStringSplit() {
	myCar := "Honda       Civic            5"
	myCarDetails := strings.Fields(myCar)
	fmt.Println(myCarDetails)
	fmt.Printf("My favorite car manufacturer is %v and my favorite model is %v, and it holds %v passengers.", myCarDetails[0], myCarDetails[1], myCarDetails[2])
}

func expandString() {
	myReplacements := [][]string{{"k8s", "Kubernetes"}, {"DNS", "Domain Name System"}, {"SNMP", "Simple Network Management Protocol"}}
	mySentence := "Hello my name is Trevor and I'm learning k8s. k8s relies on DNS heavily. Also you can monitor network devices with SNMP."

	for _, replacement := range myReplacements {
		mySentence = strings.ReplaceAll(mySentence, replacement[0], replacement[1])
	}

	fmt.Printf("\n%v\n", mySentence)
}

func testStringBuilder() {
	myBuilder := strings.Builder{}
	myBuilder.Write([]byte{235, 134, 121, 117})
	myBuilder.WriteString("Trevor Sullivan is learning Kubernetes")

	fmt.Printf("The length of builder is %v\n", myBuilder.Len())
	fmt.Printf("The capacity of builder is %v\n", myBuilder.Cap())
	fmt.Printf("The builder value is: %v", myBuilder.String())
}

func main() {
	// testFormatSpecifiers()
	// testStringSplit()
	// expandString()
	testStringBuilder()
}
