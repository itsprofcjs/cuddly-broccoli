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
	rupesh := person{
		firstName: "Rupesh Kumar",
		lastName:  "Singh",
		contactInfo: contactInfo{
			email:   "itsprofcjs@gmail.com",
			zipCode: 000000,
		},
	}

	rupesh.updateName("cjs")

	rupesh.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
