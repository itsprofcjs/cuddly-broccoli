package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {

	data := []rune{'a', 'b', 'c', 'd'}

	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))

		fmt.Printf("%c done!\n", r)
	}

	for i := 0; i < len(data); i++ {
		char := data[i]

		go capIt(char)
	}

	time.Sleep(500 * time.Microsecond)

	fmt.Printf("Capitalized: %c\n", capitalized)
}
