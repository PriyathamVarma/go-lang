// This is for explanation of functions in Go

/*
Alice and Eron have been traveling for days, and they're starting to feel tired and weak. They realize that they need to find a way to boost their energy if they want to continue their journey.

Eron suggests that they use the power of Go programming to create a function that can generate a magical potion to give them strength. He explains that functions are like recipes that we can use to create something by following a set of instructions.

To help Alice understand, Eron shows her a code example that demonstrates the concept of functions in Go. In the example, they create a function called createPotion that takes in two arguments - the ingredient name and the quantity - and uses them to create a magical potion.
*/


package main

import (
	"fmt"
)

func createPotion(ingredient string, quantity int) {
	fmt.Printf("%d ounces of %s added to cauldron\n", quantity, ingredient)
}

func main() {
	createPotion("unicorns hair", 2)
	createPotion("ravens blood", 1)
	createPotion("dragons scales", 3)
}


/*

Eron explains to Alice that functions are a powerful tool in programming that allow us to modularize our code and make it easier to manage. By creating functions for specific tasks, we can reuse code and avoid repetition.

Alice is impressed by the power of functions and is eager to learn more. Eron tells her that they can use functions to create even more complex programs and that the possibilities are endless with Go programming.

As they continue their journey, Alice and Eron create many more functions to help them on their quest. They create functions for creating powerful weapons, casting spells, and even for conjuring food and water when they need it. With their newfound knowledge of functions and Go programming, Alice and Eron become unstoppable in their quest to conquer the dark forces that threaten their way of life.


*/