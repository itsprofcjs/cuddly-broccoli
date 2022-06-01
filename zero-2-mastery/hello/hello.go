package main

import (
	"fmt"

	"rsc.io/quote"

	"cjs/greetings"
)

func main() {
	fmt.Println("Hello, CJS!")

	fmt.Println(quote.Go())

	message := greetings.Hello("CJS!")

	fmt.Println(message)
}
