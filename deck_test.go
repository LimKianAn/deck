package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52, but got %v", len(d))
	}

	correctDeck := "1 of Spades,2 of Spades,3 of Spades,4 of Spades,5 of Spades,6 of Spades,7 of Spades,8 of Spades,9 of Spades,10 of Spades,11 of Spades,12 of Spades,13 of Spades,1 of Hearts,2 of Hearts,3 of Hearts,4 of Hearts,5 of Hearts,6 of Hearts,7 of Hearts,8 of Hearts,9 of Hearts,10 of Hearts,11 of Hearts,12 of Hearts,13 of Hearts,1 of Diamonds,2 of Diamonds,3 of Diamonds,4 of Diamonds,5 of Diamonds,6 of Diamonds,7 of Diamonds,8 of Diamonds,9 of Diamonds,10 of Diamonds,11 of Diamonds,12 of Diamonds,13 of Diamonds,1 of Clubs,2 of Clubs,3 of Clubs,4 of Clubs,5 of Clubs,6 of Clubs,7 of Clubs,8 of Clubs,9 of Clubs,10 of Clubs,11 of Clubs,12 of Clubs,13 of Clubs"
	if d.toString() != correctDeck {
		t.Errorf("The contents of some cards are wrong.")
	}
}

func TestSaveToFileAndLoadDeckFromFile(t *testing.T) {
	fileName := "___temporaryFileForTesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)
	loadedDeck := loadDeckFromFile(fileName)

	correctDeck := "1 of Spades,2 of Spades,3 of Spades,4 of Spades,5 of Spades,6 of Spades,7 of Spades,8 of Spades,9 of Spades,10 of Spades,11 of Spades,12 of Spades,13 of Spades,1 of Hearts,2 of Hearts,3 of Hearts,4 of Hearts,5 of Hearts,6 of Hearts,7 of Hearts,8 of Hearts,9 of Hearts,10 of Hearts,11 of Hearts,12 of Hearts,13 of Hearts,1 of Diamonds,2 of Diamonds,3 of Diamonds,4 of Diamonds,5 of Diamonds,6 of Diamonds,7 of Diamonds,8 of Diamonds,9 of Diamonds,10 of Diamonds,11 of Diamonds,12 of Diamonds,13 of Diamonds,1 of Clubs,2 of Clubs,3 of Clubs,4 of Clubs,5 of Clubs,6 of Clubs,7 of Clubs,8 of Clubs,9 of Clubs,10 of Clubs,11 of Clubs,12 of Clubs,13 of Clubs"
	if loadedDeck.toString() != correctDeck {
		t.Errorf("The loaded deck is corrupted.")
	}

	os.Remove(fileName)
}
