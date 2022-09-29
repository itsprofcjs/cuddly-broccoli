package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(deck deck, handSize int) (deck, deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck deck) toString() string {
	return strings.Join([]string(deck), ",")
}

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(deck.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)

		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")

	return deck(stringSlice)
}

func (deck deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range deck {
		newPosition := r.Intn(len((deck)) - 1)

		deck[index], deck[newPosition] = deck[newPosition], deck[index]
	}
}
