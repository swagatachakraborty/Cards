package main

import (
	"fmt"
	"strings"
)

type card struct {
	value string
	suit string
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

func stringToCard(str string) (card) {
	suit := strings.Split(str, " of ")[0]
	value := strings.Split(str, " of ")[1]
	return card{value, suit}
}