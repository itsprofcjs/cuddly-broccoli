package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]

		fmt.Println(element)
	}
}

type Part struct {
	name string
}

func printPart(parts []Part) {

	fmt.Println()

	fmt.Println("Assembly Part")

	for i := 0; i < len(parts); i++ {
		part := parts[i]

		fmt.Println(part)
	}
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}

	printSlice("Route 1", route)

	route = append(route, "Home")

	printSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printSlice("Remaining locations", route)

	//--Summary:
	//  Create a program to manage parts on an assembly line.
	//
	//--Requirements:
	//* Using a slice, create an assembly line that contains type Part
	//* Create a function to print out the contents of the assembly line
	//* Perform the following:
	//  - Create an assembly line having any three parts
	//  - Add two new parts to the line
	//  - Slice the assembly line so it contains only the two new parts
	//  - Print out the contents of the assembly line at each step
	//--Notes:
	//* Your program output should list 3 parts, then 5 parts, then 2 parts

	assemblyLine := []Part{{name: "Head"}, {name: "Gear"}, {name: "Break"}}

	printPart(assemblyLine)

	assemblyLine = append(assemblyLine, Part{name: "Speedometer"})
	assemblyLine = append(assemblyLine, Part{name: "Chain"})

	printPart(assemblyLine)

	assemblyLine = assemblyLine[3:]

	printPart(assemblyLine)
}
