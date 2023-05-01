package person

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       uint8
}

func NewPerson() map[string]string {
	return map[string]string{"FirstName": "Trevor", "LastName": "Sullivan"}
}

func (p *Person) IncrementAge() {
	p.Age += 1
}

func (p Person) GetFullName() string {
	fullName := fmt.Sprintf("%v %v", p.FirstName, p.LastName)
	fmt.Println(fullName)
	return fullName
}
