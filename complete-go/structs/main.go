package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "21 jump street",
			zipCode: 2352252,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (person *person) updateName(newFirstName string) {
	person.firstName = newFirstName
}

func (person person) print() {
	fmt.Printf("%+v", person)
}
