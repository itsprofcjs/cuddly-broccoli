package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	deckLength := len(deck)

	if deckLength != 48 {
		t.Errorf("Expected deck length of 48, but got %v", deckLength)
	}

	if deck[0] != "Ace Of Spades" {
		t.Errorf("Expected first card of Ace of Spade but got %v", deck[0])
	}

	if deck[deckLength-1] != "King Of Clubs" {
		t.Errorf("Expected last card of King of Clubs but got %v", deck[deckLength-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"

	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	deckLength := len(loadedDeck)

	if deckLength != 48 {
		t.Errorf("Expected deck length of 48 but got %v", deckLength)
	}

	os.Remove(fileName)
}
