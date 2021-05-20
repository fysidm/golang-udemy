package main

import "fmt"

func main() {
	// basic - initial variable
	// var card string = "Ace of Spades"
	// var card string
	// card = "Ace of Spades"
	// convention - initial variable
	card := newCard()

	fmt.Println(card)

	// Array: Fixed length array
	// Slice: An array that cat grow or shrink
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// For Loop
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
