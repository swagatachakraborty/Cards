package main

import "fmt"

type card struct {
	suit string
	value string
}

func getSuits() ([]string){
	return []string{"Clubs", "Diamonds", "Hearts", "Spades"}
}

func getValues() ([]string){
	return []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Joker", "Queen", "King"}
}

func (c card) print() {
	fmt.Println(c.toString())
}

func (c card) toString() (string) {
	return (c.value + " of " + c.suit)
}