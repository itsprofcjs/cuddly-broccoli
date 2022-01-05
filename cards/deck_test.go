package main

import "testing"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	deckLength := len(deck)

	if deckLength != 48 {
		t.Errorf("Expected deck length of 48, but got %v", deckLength)
	}
}
