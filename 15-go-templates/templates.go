package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func example01() {
	t, _ := template.ParseFiles("finaloutput.txt", "block01.txt", "block02.txt")
	fmt.Println(t)
	t.Execute(os.Stdout, nil)
}

type Person struct {
	FirstName string
	LastName  string
	// Today     string
	Age uint8
}

func example02() {
	fileList := []string{"dataoutput.txt", "datablock01.txt"}
	t, _ := template.ParseFiles(fileList...)
	fmt.Println(t)
	// inputs := []string{"Daniel", "McGovern", time.Now().Format("2006-01-02"), "34"}
	// person01 := Person{FirstName: "Trevor", LastName: "Sullivan", Today: time.Now().Format("2006-01-02"), Age: 28}
	person01 := Person{FirstName: "Trevor", LastName: "Sullivan", Age: 28}
	fmt.Printf("%+v", person01)
	t.Execute(os.Stdout, person01)
}

func exampleFuncs() {
	fileList := []string{"funcs-output.txt", "funcs-block01.txt"}

	t := template.New("funcs-output.txt")

	funcList := template.FuncMap{}
	funcList["todayDate"] = func() string {
		return time.Now().Format("2006-01-02")
	}
	funcList["getYearFromString"] = func(input string) string {
		return input[0:4]
	}

	t.Funcs(funcList)
	t.ParseFiles(fileList...)
	fmt.Println(t)

	person01 := Person{FirstName: "Trevor", LastName: "Sullivan", Age: 28}
	fmt.Printf("%+v", person01)
	t.Execute(os.Stdout, person01)
}

func main() {
	// example01()
	// example02()
	exampleFuncs()
}
