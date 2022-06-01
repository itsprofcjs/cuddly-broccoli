package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

type Coordinates struct {
	x, y int
}

type Rectangle struct {
	a Coordinates
	b Coordinates
}

func length(rec Rectangle) int {
	return rec.b.y - rec.a.y
}

func breadth(rec Rectangle) int {
	return rec.b.x - rec.a.x
}

func area(rec Rectangle) int {
	return length(rec) * breadth(rec)
}

func perimeter(rec Rectangle) int {
	return 2 * (length(rec) + breadth(rec))
}

func printRectInfo(rec Rectangle) {
	fmt.Println("Area is", area(rec))
	fmt.Println("Perimeter is", perimeter(rec))
}

func main() {
	casey := Passenger{Name: "Casey", TicketNumber: 1}

	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)

	fmt.Println(bill, ella)

	var heidi Passenger

	heidi.Name = "Heidi"
	heidi.TicketNumber = 4

	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println(bill.Name, " as boarded the bus")
	}

	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.Boarded = true

	bus := Bus{FrontSeat: heidi}

	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

	//--Summary:
	//  Create a program to calculate the area and perimeter
	//  of a rectangle.
	//
	//--Requirements:
	//* Create a rectangle structure containing its coordinates
	//* Using functions, calculate the area and perimeter of a rectangle,
	//  - Print the results to the terminal
	//  - The functions must use the rectangle structure as the function parameter
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal
	//
	//--Notes:
	//* The area of a rectangle is length*width
	//* The perimeter of a rectangle is the sum of the lengths of all sides

	rec := Rectangle{a: Coordinates{x: 2, y: 5}, b: Coordinates{x: 10, y: 15}}

	fmt.Println("****************************** Calculating the area ***************************")

	printRectInfo(rec)

	rec.a.y *= 2
	rec.b.y *= 2

	fmt.Println("****************************** Doubling the area ******************************")

	printRectInfo(rec)

}
