package main

import (
	"io/ioutil"
)

type deck []card

func getNewDeck() (deck) {
	suits := getSuits()
	values := getValues()

	return createDeck(suits, values)
}

func createDeck(suits []string, values []string) (deck) {
	newDeck := deck{}
	for _, suit := range suits {
		for _, value := range values {
			newDeck = append(newDeck, card{suit, value})
		}
	}
	return newDeck
}

func (d deck) print() {
	for _, c := range d {
		c.print()
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() (string) {
	var stringifyDeck string

	for _, c := range []card (d) {
		
		if len(stringifyDeck) > 0 {
			stringifyDeck += "\n"
		}
		
		stringifyDeck += c.toString()
	}

	return stringifyDeck
}

func (d deck) saveTheDeck(filepath string) error {
	return ioutil.WriteFile(filepath, []byte (d.toString()), 0666)
}