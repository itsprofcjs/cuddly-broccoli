package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()

	fmt.Println("------------")

	hand, remainingDeck := deal(cards, 15)

	hand.print()

	fmt.Println("------------")

	remainingDeck.print()

	fmt.Println("------------")

	fmt.Println(cards.toString())

	fmt.Println("******** saving cards ********")

	cards.saveToFile("last_cards")

	fmt.Println("******** reading cards **********")

	savedCards := newDeckFromFile("last_cards")

	savedCards.print()

	fmt.Println("********* Suffle cards **********")

	savedCards.shuffle()

	savedCards.print()
}
