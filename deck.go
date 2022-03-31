package main

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
