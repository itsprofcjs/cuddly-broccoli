package main

import "fmt"

func double(x int) int {
	return 2 * x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from greet function")
}

func greetPerson(name string) {
	fmt.Println("Hello ", name)
}

func stringReturnFunc() string {
	return "This is string return function"
}

func add3Numbers(firstNumber, secondNumber, thirdNumber int) int {
	return firstNumber + secondNumber + thirdNumber
}

func returnAnyNumber() int {
	return 6
}

func returnTwoNumber() (int, int) {
	return 3, 6
}

func main() {

	greet()

	doubleNumber := double(6)
	fmt.Println("A double is", doubleNumber)

	doubleAdded := add(doubleNumber, 10)
	fmt.Println("10 added doubleNumber ", doubleAdded)

	//--Summary:
	//  Use functions to perform basic operations and print some
	//  information to the terminal.
	//
	//--Requirements:
	//* Write a function that accepts a person's name as a function
	//  parameter and displays a greeting to that person.
	greetPerson("CJS")
	//* Write a function that returns any message, and call it from within
	//  fmt.Println()
	fmt.Println("stringReturnFunc", stringReturnFunc())
	//* Write a function to add 3 numbers together, supplied as arguments, and
	//  return the answer
	fmt.Println("add3Numbers", add3Numbers(1, 2, 3))
	//* Write a function that returns any number
	fmt.Println("returnAnyNumber", returnAnyNumber())
	//* Write a function that returns any two numbers
	oneNumber, twoNumber := returnTwoNumber()
	fmt.Println("returnTwoNumber", oneNumber, twoNumber)
	//* Add three numbers together using any combination of the existing functions.
	fmt.Println("Add 3 number", add3Numbers(oneNumber, returnAnyNumber(), 1))
	//  * Print the result
	//* Call every function at least once

}
