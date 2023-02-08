// Packages
package main

// Imports
import (
	"log"
	"time"
)

// Functions
func main() {

	//cards := []string{newCard(), newCard()} // declaring an array
	// cards := deck{newCard(), newCard()}
	// cards = append(cards, "New record") // pushing values to an array
	// Timing the code
	start := time.Now()

	cards := newDeck()
	cards.print()

	elapsed := time.Since(start)
	log.Printf("Time taken to execute", elapsed)

}

func newCard() string {
	return "Arrays Concept"
}
