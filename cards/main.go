package main

import "fmt"

func main() {
	// basic - initial variable
	// var card string = "Ace of Spades"
	// convention - initial variable
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
