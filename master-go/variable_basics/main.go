package main

import "fmt"

func main() {
	name := "Dan"
	// unused variables could be used as
	_ = name

	fmt.Println("variables", name)
}
