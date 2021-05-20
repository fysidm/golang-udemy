package main

func main() {
	// basic - initial variable
	// var card string = "Ace of Spades"
	// var card string
	// card = "Ace of Spades"
	// convention - initial variable
	// card := newCard()

	// Array: Fixed length array
	// Slice: An array that cat grow or shrink
	cards := newDeck()

	// For Loop
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	cards.shuffle()
	cards.print()
}
