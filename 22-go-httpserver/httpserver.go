package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	SimpleHttpServer()
}

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       uint8  `json:"Age"`
}

var People []Person = []Person{
	{FirstName: "John", LastName: "McGovern", Age: 30},
}

type Dog struct {
	Name  string `json:"Name"`
	Breed string `json:"Breed"`
}

var DogData []Dog = []Dog{
	{Name: "Cheyenne", Breed: "Border Collie"},
	{Name: "Rover", Breed: "Great Dane"},
}

func DogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dogjson, _ := json.Marshal(DogData)
		w.Write(dogjson)
	}
}

func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, _ := json.Marshal(People)
		// data := "Hello from Go REST API"
		w.Write(data)
	}
	if r.Method == "PUT" {
		data, _ := io.ReadAll(r.Body)
		newPerson := Person{}
		json.Unmarshal(data, &newPerson)
		People = append(People, newPerson)
		w.WriteHeader(http.StatusCreated)
	}
	if r.Method == "DELETE" {
		data, _ := io.ReadAll(r.Body)
		newPerson := Person{}
		json.Unmarshal(data, &newPerson)

		for i, person := range People {
			if person.FirstName == newPerson.FirstName && person.LastName == newPerson.LastName {
				People = append(People[:i], People[i+1])
			}
		}
	}
}

func SimpleHttpServer() {

	dogmux := http.NewServeMux()
	dogmux.HandleFunc("/api/v1/dog", DogHandler)
	go http.ListenAndServe(":9544", dogmux)

	http.HandleFunc("/api/v1/people", PeopleHandler)
	fmt.Println("Starting HTTP server on port 8074")
	http.ListenAndServe(":8074", nil)
}
