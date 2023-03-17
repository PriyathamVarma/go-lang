// This is to explain the concept of switch statements in Go

/*

Alice and Eron were now experts in Go programming, but they knew there was always more to learn. They decided to use their knowledge to create a program that would help them find the perfect ingredients for their magic potion.

To do this, they used a conditional statement called a switch statement in their program. A switch statement is used to evaluate an expression and execute different code blocks depending on the value of the expression.

In their program, Alice and Eron created a switch statement that would evaluate the user's input and determine which ingredient to add to the potion based on their selection. Here is an example code snippet that demonstrates the use of switch statement:

*/

package main

import (
	"fmt"
)

func main() {
	var ingredient string

	fmt.Print("Enter an ingredient: ")
	fmt.Scanln(&ingredient)

	switch ingredient {
	case "unicorns hair":
		fmt.Println("Adding unicorns hair to the potion.")
	case "ravens blood":
		fmt.Println("Adding ravens blood to the potion.")
	case "dragons scales":
		fmt.Println("Adding dragons scales to the potion.")
	case "mermaid tears":
		fmt.Println("Adding mermaid tears to the potion.")
	default:
		fmt.Println("Invalid ingredient.")
	}
}


/*

In this example, the user is prompted to enter an ingredient. The input is stored in the ingredient variable. The switch statement then evaluates the value of ingredient and executes the corresponding code block.

If the user enters "unicorns hair", the program adds unicorns hair to the potion. If the user enters "ravens blood", the program adds ravens blood to the potion, and so on. If the user enters an invalid ingredient, the program displays a message saying "Invalid ingredient."

Through the use of switch statements, Alice and Eron were able to create a program that helped them find the perfect ingredients for their magic potion. They were amazed at how versatile and powerful Go programming was, and they knew that they had only scratched the surface of what was possible with this incredible language.





*/