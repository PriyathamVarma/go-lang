// This is to explain the concept of variables in Go

/*

As Alice and Eron continued their journey, they came across a problem. They needed to create a powerful wand that could protect them against the dark forces that threatened their way of life.

Eron told Alice that they could use Go programming to create the perfect wand. He explained that Go programming uses variables to store and manipulate data in programs.

To help Alice understand, Eron showed her a code example that demonstrated the concept of variables in Go. In the example, they declared different types of variables such as string, int, and bool, and assigned values to them using different syntaxes.

Eron explained to Alice that a variable is like a container that holds a value. For example, a variable called name can hold the value "Alice", and a variable called age can hold the value "20". By assigning values to variables, we can store data and use it later in the program.

Eron also taught Alice that variables can be declared and assigned values in different ways, depending on the situation. For example, the short variable declaration := can be used to declare and assign a value to a variable in one step.

Alice was fascinated by the concept of variables and how they could be used to create powerful programs. She asked Eron how variables could be used to create the perfect wand.

Eron smiled and said, "Well, we can use variables to store the wand's properties, such as its length, width, and material. We can then use these variables to calculate the wand's power and make adjustments as needed. With Go programming, we have the flexibility to create the perfect wand for any situation."

Alice was excited by the possibilities and eager to continue learning about Go programming and the power of Verseium. She knew that with her newfound knowledge of variables and programming, she could create anything she could imagine.

*/

package main

import (
	"fmt"
)

func main() {

	// This is a variable declaration
	var nameofTheWand string

	// This is an assignment statement
	nameofTheWand = "Alices wonderful wand"

	// This is a short variable declaration
	ageofTheWand := 20

	// This is a variable declaration with an initial value
	var isChanted bool = true

	// This is a variable declaration with multiple variables
	var (
		wandHandle string  = "wood"
		wandMantle string = "Verseium"
	)

	fmt.Println(nameofTheWand, ageofTheWand, isChanted, wandHandle, wandMantle);

}

/* 

In continuation of Alice and Eron's quest, they decided to create a powerful wand using Go programming and the concept of variables.

To create the perfect wand, they needed to store the wand's properties, such as its name, age, handle, and mantle, in variables.

Eron showed Alice how to declare and assign values to variables in Go. In the code example, they declared different types of variables such as string, int, and bool, and assigned values to them using different syntaxes.

Alice was excited to see how variables could be used to store the properties of the wand. They declared variables like nameofTheWand, ageofTheWand, isChanted, wandHandle, and wandMantle and assigned values to them to represent the properties of the wand.

Through this example, Alice learned that variables are important in programming because they allow us to store and manipulate data in our programs. By using variables, we can make our programs more flexible and dynamic, and we can write code that can adapt to different situations.

Finally, Eron and Alice printed the values of the variables to see the properties of the wand they created. Alice was thrilled to see the properties of the wand in the output and knew that they were one step closer to creating the perfect wand.

With this newfound knowledge of variables and programming, Alice and Eron continued their quest to create the most powerful wand that could protect them against any dark force that threatened their way of life.

*/