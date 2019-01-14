package main

import (
	"os"
	"strconv"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but instead got %v", len(d))
	}

	expected := deck([]string{"Ace of Spades", "Two of Spades", "Three of Spades", "Four of Spades", "Five of Spades", "Six of Spades", "Seven of Spades", "Eight of Spades", "Nine of Spades", "Ten of Spades", "Jack of Spades", "Queen of Spades", "King of Spades", "Ace of Diamonds", "Two of Diamonds", "Three of Diamonds", "Four of Diamonds", "Five of Diamonds", "Six of Diamonds", "Seven of Diamonds", "Eight of Diamonds", "Nine of Diamonds", "Ten of Diamonds", "Jack of Diamonds", "Queen of Diamonds", "King of Diamonds", "Ace of Hearts", "Two of Hearts", "Three of Hearts", "Four of Hearts", "Five of Hearts", "Six of Hearts", "Seven of Hearts", "Eight of Hearts", "Nine of Hearts", "Ten of Hearts", "Jack of Hearts", "Queen of Hearts", "King of Hearts", "Ace of Clubs", "Two of Clubs", "Three of Clubs", "Four of Clubs", "Five of Clubs", "Six of Clubs", "Seven of Clubs", "Eight of Clubs", "Nine of Clubs", "Ten of Clubs", "Jack of Clubs", "Queen of Clubs", "King of Clubs"})
	for i, card := range expected {
		if card != d[i] {
			t.Errorf("Expected " + card + " at index " + strconv.Itoa(i) + " of deck, instead got " + d[i])
		}
	}
}

func TestSaveToFileAndLoadDeckFromFile(t *testing.T) {
	//  File name
	fileName := "_deckTesting.txt"

	// Clean up file from previous test
	os.Remove(fileName)

	// Instiate a new deck
	deck := newDeck()

	// Save the deck to file
	deck.saveToFile(fileName)

	// Load the deck from file
	loadedDeck := loadDeckFromFile(fileName)

	// Make sure the loaded deck is the correct length
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	// Make sure the saved deck and loaded deck are the same
	for i, card := range loadedDeck {
		if card != deck[i] {
			t.Errorf("Expected " + card + " at index " + strconv.Itoa(i) + " of deck, instead got " + deck[i])
		}
	}

	// Clean up after test
	os.Remove(fileName)

}
