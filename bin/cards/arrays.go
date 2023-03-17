// Packages
package main

// Imports
import "fmt"

// Functions
func main() {

	cards := []string{newCard(), newCard()} // declaring an array
	cards = append(cards, "New record")     // pushing values to an array

	for i, card := range cards {

		fmt.Println(i, card)

	}

}

func newCard() string {
	return "Arrays Concept"
}
