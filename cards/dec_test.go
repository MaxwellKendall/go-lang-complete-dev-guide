package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %v", len(d))
	}

	// first card is ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}
	// last card is the king of diamonds
	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds but got %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck, _ := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck to be of length 52 instead got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}