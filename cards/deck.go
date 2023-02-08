// Packages
package main

import "fmt"

// imports

// create a new type from basic types
type deck []string

// This type can be used anywhere in the files

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clues", "Flowers"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	cardExtras := []string{"Test1", "Test2", "Test3", "test4", "test5"} // this i have written

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			for _, extra := range cardExtras {
				cards = append(cards, suit+"of"+value+"which is"+extra)
			}
		}
	}

	return cards
}

// functions
func (d deck) print() { // here d deck is a reciever

	for i, card := range d {

		fmt.Println(i, card)

	}

}
