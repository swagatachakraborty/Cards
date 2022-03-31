package main

import "fmt"

func main() {
	fmt.Println("Creating new deck of cards:")
	fmt.Println("---------------------------")
	cards := getNewDeck()
	cards.print()

	fmt.Println("\nShuffling the deck")
	fmt.Println("------------------")
	cards.shuffle()
	cards.print()

	fmt.Println("\nMaking a deal of 5 cards")
	fmt.Println("------------------------")
	deal, remainingDeck := cards.deal(5)
	fmt.Println("Deal:")
	fmt.Println("-----")
	deal.print()
	fmt.Println("Remaining Deck:")
	fmt.Println("---------------")
	remainingDeck.print()

	fmt.Println("\nSaving deal to file my_deal")
	fmt.Println("---------------------------")
	deal.saveTheDeck("my_deal")

	fmt.Println("\nFetching saved deal from file my_deal")
	fmt.Println("-------------------------------------")
	fetchSavedDeckFromFile("my_deal").print()
}
