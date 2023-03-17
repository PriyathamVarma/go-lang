// This is to explain the concept of loops in a different way in Go

package main

import (
"fmt"
)

func main() {

	/*
	Alice and Eron are on a quest to find a rare flower that can cure a mysterious illness that has been afflicting the people of their village. They know that the flower only blooms in a certain area of the forest and that they will need to search through many patches of flowers to find it.

	They start by creating a slice of flowers that they have already searched through and another slice of flowers that they have yet to search through.
*/

searchedFlowers := []string{"daffodil", "tulip", "daisy", "lily"}
unsearchedFlowers := []string{"rose", "carnation", "poppy", "orchid", "snapdragon", "petunia"}

/*
	Alice and Eron decide to use a loop to search through each patch of flowers. They use a for loop that iterates through each item in the unsearchedFlowers slice and checks if it is the rare flower they are looking for.

	Once they find the rare flower, they break out of the loop and proceed to make the cure for the illness.
*/

rareFlower := "snapdragon"
foundFlower := false

for _, flower := range unsearchedFlowers {
	if flower == rareFlower {
		foundFlower = true
		break
	} else {
		searchedFlowers = append(searchedFlowers, flower)
	}
}

/*
	Alice and Eron check if they found the rare flower. If they did, they proceed to make the cure using a separate function. If not, they continue searching through the remaining patches of flowers.
*/

if foundFlower {
	fmt.Println("We found the rare flower! Let's make the cure.")
	makeCure()
} else {
	fmt.Println("We didn't find the rare flower in this patch. Let's move on to the next one.")
}

/*
	Through this example, Alice and Eron learn that loops are useful for repeating a set of instructions until a certain condition is met. They can use loops to search through large sets of data or perform repetitive tasks with ease.

	With their newfound knowledge of loops, Alice and Eron are able to successfully find the rare flower and cure the mysterious illness that had been afflicting their village.
*/

}

func makeCure() {
fmt.Println("The cure has been made!")
}


}

