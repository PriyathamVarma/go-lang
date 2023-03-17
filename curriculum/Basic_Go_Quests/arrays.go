// This is to explain the concept of arrays in Go

/*

As Alice and Eron continued their quest to create the most powerful wand, they came across the concept of arrays in Go programming.

Eron explained to Alice that an array is a collection of elements of the same type that are stored together in a contiguous block of memory. Arrays are useful for storing and accessing a fixed number of elements.

To help Alice understand, Eron showed her a code example that demonstrated the concept of arrays in Go. In the example, they declared arrays of different types, such as int, bool, and string, and assigned values to them using different syntaxes.

Eron explained to Alice that in Go, arrays are zero-indexed, which means that the first element in an array has an index of 0, the second element has an index of 1, and so on.

Alice was fascinated by the concept of arrays and how they could be used to store multiple values of the same type. They declared variables like wandLength, wandWidth, isChanted, wandHandle, and wandMantle as arrays and assigned values to them to represent the properties of the wand.

Through this example, Alice learned that arrays are useful for storing and accessing a fixed number of elements of the same type in a program. By using arrays, we can store and manipulate data in our programs in an efficient manner.

*/

package main

import (

"fmt"

)

func main() {

// Dynamic array

var wandlengthOptions [...]int;

// This is a variable declaration

var wandLengths [3]int

// This is an assignment statement

wandLengths[0] = 10

wandLengths[1] = 20

wandLengths[2] = 30

// This is a short variable declaration

wandWidths := [3]int{10, 20, 30}

// This is a variable declaration with an initial value

var isChanted [3]bool = [3]bool{true, false, true}

// This is a variable declaration with multiple variables

var (

wandHandles [3]string = [3]string{"wood", "glass", "metal"}

wandMantle [3]string = [3]string{"Verseium", "Sleseium", "Demiseium"}// Sleseium and Demiseium are anti Verseium and can be used to create a wand that can destroy Verseium wands

)

fmt.Println(wandLengths, wandWidths, isChanted, wandHandles, wandMantles);

}

/* 

Finally, Eron and Alice printed the values of the arrays to see the properties of the wand they created. Alice was thrilled to see the properties of the wand stored in arrays and knew that arrays would be useful in creating the perfect wand that could protect them against any dark force that threatened their way of life.

With this newfound knowledge of arrays and programming, Alice and Eron continued their quest to create the most powerful wand that could protect them against any dark force that threatened their way of life.

*/