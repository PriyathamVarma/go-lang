// This is to explain the concept of slices in Go

package main

import (
"fmt"
)

func main() {

	/* Alice and Eron are trying to create a list of ingredients for a potion. They start by creating an array to store the ingredients:*/
// Create an array with the ingredients for the potion

var ingredients []string = []string{"unicorns hair","ravens blood","dragons scales","mermaid tears"}


// However, they realize that they only need the first three ingredients for their potion. They can create a slice of the first three ingredients like this:
// Create a magic portion with the ingredients

magicPortion := ingredients[0:3];// ->>> "unicorns hair","ravens blood","dragons scales"


/*
This creates a slice magicPortion that includes the first three elements of the ingredients array.

Later, Alice and Eron discover a new ingredient, "phoenix feathers", that they want to add to their potion. They can add this ingredient to their slice using the append function:
*/
// To append new ingredients to the magic portion

magicPortion = append(magicPortion, "phoenix feathers");

fmt.Println(magicPortion);// ->>>> "unicorns hair","ravens blood","dragons scales","phoenix feathers"	


}

/*

This adds the value "phoenix feathers" to the end of the magicPortion slice.

Through this example, Alice and Eron learn that slices are useful for working with subsets of an array in a flexible manner. They can create a "window" into an array by using slices, and then modify the elements within that window as needed.

With their newfound knowledge of slices, Alice and Eron are able to create the perfect potion using just the ingredients they need.

*/
