package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	// contactInfo contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPersion *person) updateName(newFirstName string) {
	(*pointerToPersion).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
