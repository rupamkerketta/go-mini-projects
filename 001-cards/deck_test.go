package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
    d := newDeck()

    if len(d) != 52 {
        t.Errorf("Expected deck length of 52, but got %v", len(d))
    }

    if d[0] != "Ace of Spades" {
        t.Errorf("Expected Ace of Spades, but got %v", d[0])
    }

    if d[len(d) - 1] != "King of Diamonds" {
        t.Errorf("Expected King of Diamonds, but got %v", d[len(d) -1])
    }
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
    // 1. Remove a file if it exists with the name of '_decktesting'
    os.Remove("_decktesting")

    // 2. Create a new deck and save it to a file named '_deckstring'
    d := newDeck()
    d.saveToFile("_decktesting")

    // 3. Load the saved deck from the file '_decktesting'
    loadedDeck := newDeckFromFile("_decktesting")

    // 4. Check the size of the deck - should be equal to 52
    if len(loadedDeck) != 52 {
        t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
    }

    // 5. Remove the file '_decktesting'
    os.Remove("_decktesting")
}
