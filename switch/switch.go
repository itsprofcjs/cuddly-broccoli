package main

import "fmt"

func price() int {
	return 3
}

const (
	Economy    = 2
	Business   = 4
	FirstClass = 6
)

func main() {

	switch p := price(); {
	case p <= Economy:
		fmt.Println("Cheap Class")
	case p <= Business:
		fmt.Println("Good Class")
		fallthrough
	default:
		fmt.Println("First Class")
	}

	//--Summary:
	//  Create a program to display a classification based on age.
	//
	//--Requirements:
	//* Use a `switch` statement to print the following:
	//  - "newborn" when age is 0
	//  - "toddler" when age is 1, 2, or 3
	//  - "child" when age is 4 through 12
	//  - "teenager" when age is 13 through 17
	//  - "adult" when age is 18+

	switch age := 2; {
	case age == 0:
		fmt.Println("newborn")
	case age < 4:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}
