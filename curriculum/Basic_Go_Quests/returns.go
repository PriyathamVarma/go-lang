// This is to explain the concept of function returns in a different way in Go

/*

Alice and Eron have been busy creating magical potions with their newfound knowledge of Go programming. They've used functions to create various steps in their potion-making process, but now they want to take things to the next level.

One day, they encounter a mysterious potion that needs to be reverse-engineered. The potion is made up of multiple ingredients, and each ingredient needs to be analyzed in a different way. They quickly realize that they need to create a function that can analyze each ingredient and return the result.

To create this function, they start by declaring the function name and the input parameter(s) it needs, as well as the type of data the function will return. In this case, they'll create a function called analyzeIngredient that takes in a string ingredient and returns a string:


*/

package main

import (

"fmt"

)

// Analyze the ingrediet and return the result
func analyzeIngredient(ingredient string) string {
    switch ingredient {
    case "unicorn hair":
        return "This ingredient has magical properties that make it perfect for potions that need a boost of power."
    case "raven's blood":
        return "This ingredient is a potent blood that's great for dark magic spells."
    case "dragon scales":
        return "These scales are incredibly durable and can be used to create powerful defensive potions."
    case "mermaid tears":
        return "Mermaid tears have healing properties and can be used to cure a variety of ailments."
    default:
        return "Sorry, we don't know how to analyze this ingredient yet."
    }
}


/*

With the function complete, Alice and Eron can now use it to analyze each ingredient in the mysterious potion and determine its properties. For example:

*/

func main() {

	potions := []string{"unicorn hair", "raven's blood", "dragon scales", "mermaid tears", "unknown ingredient"}//slice of strings

	for _, ingredient := range potions {
		result := analyzeIngredient(ingredient)
		fmt.Println(result)
	}

}

/*

This code will loop through the list of ingredients in the potions slice, call the analyzeIngredient function on each one, and print out the result. For the first four ingredients, the function will return a string describing their properties. For the last ingredient, which is unknown, the function will return a default message.

Through this example, Alice and Eron learn the concept of functions that return a value, which can be incredibly useful in analyzing data and making decisions in their potion-making process.

*/