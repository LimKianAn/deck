package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type deck []string

func newDeck() deck {
	dck := deck{}
	cardSuits := [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	for _, suit := range cardSuits {
		for index := range [13]int{} {
			card := strconv.Itoa(index+1) + " of " + suit
			dck = append(dck, card)
		}
	}

	return dck
}

func loadDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))
}

func (dck deck) shuffle() {
	newStrings := make([]string, len(dck))
	for i, element := range newRandomSequence(len(dck)) {
		newStrings[i] = dck[element-1]
	}

	for i := range dck {
		dck[i] = newStrings[i]
	}
}

func (dck deck) print() {
	for index, card := range dck {
		fmt.Println(index, card)
	}
}

func (dck deck) toString() string {
	return strings.Join(dck, ",")
}

func (dck deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(dck.toString()), 0666)
}

func deal(dck deck, handSize int) (deck, deck) {
	return dck[:handSize], dck[handSize:]
}
