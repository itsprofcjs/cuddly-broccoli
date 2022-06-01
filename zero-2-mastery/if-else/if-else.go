package main

import "fmt"

func average(firstNumber, secondNumber, thirdNumber int) int {
	return (firstNumber + secondNumber + thirdNumber) / 3
}

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func weekdays(day int) bool {
	return day < 5
}

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func main() {
	quiz1, quiz2, quiz3 := 3, 2, 3

	if quiz1 > quiz2 {
		fmt.Println("quiz 1 scored higher than quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz 2 scored higher than quiz 1")
	} else {
		fmt.Println("quiz 1 & quiz 2 have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable")
	} else {
		fmt.Println("Not Acceptable")
	}

	//--Summary:
	//  The existing program is used to restrict access to a resource
	//  based on day of the week and user role. Currently, the program
	//  allows any user to access the resource. Implement the functionality
	//  needed to grant resource access using any combination of `if`, `else if`,
	//  and `else`.
	//
	//--Requirements:
	//* Use the accessGranted() and accessDenied() functions to display
	//  informational messages
	//* Access at any time: Admin, Manager
	//* Access weekends: Contractor
	//* Access weekdays: Member
	//* Access Mondays, Wednesdays, and Fridays: Guest

	today, role := Saturday, Member

	var hasAccess = false

	if role == Admin || role == Manager {
		hasAccess = true
	} else if role == Contractor {
		if !weekdays(today) {
			hasAccess = true
		}
	} else if role == Member {
		if weekdays(today) {
			hasAccess = true
		}
	} else if role == Guest {
		if today == Monday || today == Wednesday || today == Friday {
			hasAccess = true
		}
	}

	if hasAccess {
		accessGranted()
	} else {
		accessDenied()
	}

}
