package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCard := deal(cards, 5)

	hand.print()
	remainingCard.print()

	fmt.Println(cards.toString())

	cards.saveToFile("cards.txt")

	cards = newDeckFromFile("cards.txt")

	cards.shuffle()

	cards.print()
}
