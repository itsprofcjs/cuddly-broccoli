package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is not clean")
		}
	}
}

type Shopping struct {
	name  string
	price int
}

func printShoppingDetail(items [4]Shopping) {
	var totalItems, totalCost int

	for i := 0; i < len(items); i++ {
		item := items[i]

		totalCost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item", items[totalItems-1])
	fmt.Println("Total number items", totalItems)
	fmt.Println("Total cost", totalCost)
}

func main() {

	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)

	fmt.Println("************************ Performing cleanup ************************")

	rooms[0].cleaned = true
	rooms[2].cleaned = true

	checkCleanliness(rooms)

	//--Summary:
	//  Create a program that can store a shopping list and print
	//  out information about the list.
	//
	//--Requirements:
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//  - Products must include the price and the name
	//* Insert 3 products into the array
	//* Print to the terminal:
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	//* Add a fourth product to the list and print out the
	//  information again

	items := [4]Shopping{
		{name: "item1", price: 10},
		{name: "item2", price: 20},
		{name: "item3", price: 30},
	}

	printShoppingDetail(items)

	fmt.Println("adding the fourth item")

	items[3].name = "item4"
	items[3].price = 40

	printShoppingDetail(items)
}
