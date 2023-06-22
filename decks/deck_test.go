package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expect deck length of 16, but got %v", len(d))
	}

	firstCardExpect := "Ace of Spades"
	if d[0] != firstCardExpect {
		t.Errorf("Expect first card of %v, but got %v", firstCardExpect, d[0])
	}

	lastCardExpect := "Four of Clubs"
	if d[len(d)-1] != lastCardExpect {
		t.Errorf("Expect last card of %v, but got %v", lastCardExpect, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deckNameTest := "_deckTesting"
	os.Remove(deckNameTest)

	deck := newDeck()

	deck.saveToFile(deckNameTest)
	loadedDeck := newDeckFromFile(deckNameTest)

	deckLengthExpected := 16
	deckLengthReturns := len(loadedDeck)
	if deckLengthReturns != deckLengthExpected {
		t.Errorf("Expected %v cards in deck, got %v", deckLengthExpected, deckLengthReturns)
	}

}
