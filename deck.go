package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

func getNewDeck() (deck) {
	suits := getSuits()
	values := getValues()
	
	return createDeck(suits, values)
}

func createDeck(suits []string, values []string) (deck) {
	newDeck := deck{}
	for _, s := range suits {
		for _, v := range values {
			newDeck = append(newDeck, card{v, s})
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

func stringToDeck(str string) deck {
	var d deck
	for _, stringifyCards := range strings.Split(str, "\n") {
		d = append(d, stringToCard(stringifyCards))
	}
	return d
}

func (d deck) saveTheDeck(filepath string) error {
	return ioutil.WriteFile(filepath, []byte (d.toString()), 0666)
}

func fetchSavedDeckFromFile(filepath string) (deck) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return stringToDeck(string (b))
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}