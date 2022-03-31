package main

import "fmt"

func main() {
	fmt.Println("Creating new deck of cards: ")
	cards := getNewDeck()
	cards.print()

	fmt.Println("------------------------------------")
	fmt.Println("Making a deal of 5 cards")
	deal, remainingDeck := cards.deal(5)
	fmt.Println("Deal:")
	deal.print()
	fmt.Println("Remaining Deck:")
	remainingDeck.print()
}	
