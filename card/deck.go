package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

// create a new type of 'deck'
// which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	
	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
	
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck)saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}






