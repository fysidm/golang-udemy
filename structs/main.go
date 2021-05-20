package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	// Updating struct values
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	// Log out alex
	fmt.Println(alex)
	// Check alex with struct key
	fmt.Printf("%+v", alex)
}
