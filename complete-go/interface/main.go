package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hi CJS!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(bot bot) {
	fmt.Println(bot.getGreeting())
}
