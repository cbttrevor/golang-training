package person

import (
	"fmt"
	"testing"
)

func Example() {
	person01 := Person{FirstName: "Trevor", LastName: "Sullivan", Age: 30}
	person02 := Person{FirstName: "John", LastName: "McGovern", Age: 28}
	fmt.Printf("%v %v\n", person02.FirstName, person02.LastName)
	fmt.Printf("%v %v\n", person01.FirstName, person01.LastName)
	// Unordered output: Trevor Sullivan
	// John McGovern
}

func TestNewPerson(t *testing.T) {
	person01 := NewPerson()

	t.Run("PersonFirstName", func(t *testing.T) {
		t.Log("The first name is being tested")
		if person01["FirstName"] != "Trevor" {
			t.Log("The first name did not match")
			t.Fail()
		}
	})

	t.Run("PersonLastName", func(t *testing.T) {
		t.Log("The last name is being tested")
		if person01["LastName"] != "Sullivan" {
			t.Log("The last name did not match")
			t.Fail()
		}
	})

}
