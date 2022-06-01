package main

import (
	"fmt"
)

func main() {
	// deduced type
	var myName = "CJS! - MyName"
	fmt.Println("My Name is", myName, myName)

	// explicit type annotation
	var name string = "Rupesh"
	fmt.Println("name = ", name)

	// create and assign
	userName := "Kumar"
	fmt.Println("userName =", userName)

	// declaration
	var sum int
	fmt.Println("The sum is", sum)

	// compound assignment
	part1, part2 := 1, 5
	fmt.Println("part 1 is", part1, "part 2 is", part2)

	// Ok
	part3, part2 := 2, 0
	fmt.Println("part 3 is", part3, "part 2 is", part2)

	sum = part1 + part3
	fmt.Println("Sum = ", sum)

	sumPart12 := part1 + part2
	fmt.Println("sum part1 + part2", sumPart12)

	sumPart32 := part3 + part2
	fmt.Println("sum part3 + part2", sumPart32)

	// block assignment
	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("LessonName=", lessonName)
	fmt.Println("LessonType", lessonType)

	// ignore compound assignment
	world1, world2, _ := "Hello", "world", "!"
	fmt.Println(world1, world2)

	//Summary:
	//  Print basic information to the terminal using various variable
	//  creation techniques. The information may be printed using any
	//  formatting you like.
	//
	//Requirements:
	//* Store your favorite color in a variable using the `var` keyword
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	//* Store your first & last initials in two variables using block assignment
	//* Declare (but don't assign!) a variable for your age (in days),
	//  then assign it on the next line by multiplying 365 with the age
	// 	variable created earlier
	//
	//Notes:
	//* Use fmt.Println() to print out information
	//* Basic math operations are:
	//    Subtraction    -
	// 	  Addition       +
	// 	  Multiplication *
	// 	  Division       /

	var color = "Red"
	fmt.Println("Variables color", color)

	// compound assignment
	year, age := 1993, 29
	fmt.Println("year", year, "age", age)

	// block assignment
	var (
		firstName = "Rupesh"
		lastName  = "Singh"
	)
	fmt.Println("firstName", firstName, "lastName", lastName)

	var ageInDays int
	ageInDays = age * 365
	fmt.Println("ageInDays", ageInDays)
}
