package main

import (
	"fmt"
	"itsprofcjs/greetings"
	"log"
)

func main() {
	//
	log.SetPrefix("greetings: single")
	// log.SetFlags(0)

	greet("CJS")
	greets(nil)
}

func greet(name string) {

	message, error := greetings.Hello(name)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}

func greets(names []string) {

	nameList := []string{"CJS", "Gladys", "Samantha", "Darrin"}

	if names != nil {
		nameList = names
	}

	messages, error := greetings.Hellos(nameList)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
