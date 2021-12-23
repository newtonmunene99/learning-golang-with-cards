package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckLength := len(d)
	expectedDeckLength := 52

	if deckLength != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, deckLength)
	}

	firstCard := d[0]
	expectedFirstCard := "Ace of Spades"

	if firstCard != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, firstCard)
	}

	lastCard := d[len(d)-1]
	expectedLastCard := "King of Clubs"

	if lastCard != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, lastCard)
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck, _ := newDeckFromFile(filename)

	deckLength := len(loadedDeck)
	expectedDeckLength := 52

	if deckLength != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, deckLength)
	}

	os.Remove(filename)
}
