package main

import "fmt"

// Create a new type of "deck" which is a slice of strings
type deck []string

// Any variable of type "deck" now gets access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
