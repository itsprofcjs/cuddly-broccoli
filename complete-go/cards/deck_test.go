package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	deckLength := len(deck)

	if deckLength != 52 {
		t.Errorf("Expected deck length of 52, but got %v", deckLength)
	}
}

func TestFirstCard(t *testing.T) {
	deck := newDeck()

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected fist card of Ace of Spades, but got %v", deck[0])
	}
}

func TestLastCard(t *testing.T) {
	deck := newDeck()

	lastCard := deck[len(deck)-1]

	if lastCard != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", lastCard)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting.txt"

	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	loadedDeckLength := len(loadedDeck)

	if loadedDeckLength != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", loadedDeckLength)
	}

	os.Remove(fileName)
}
