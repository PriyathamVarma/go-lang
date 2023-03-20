// This is to explain the concept of maps in Go

package main

import "fmt"

/*

Alice and Eron are now experts in creating magical potions and analyzing ingredients using functions and structs. However, they face a new challenge. They need to keep track of the ingredients they have and the quantities they possess. They realize that a map data structure would be perfect for this task.

A map in Go is a collection of key-value pairs, where each key is unique and the value can be of any type. Alice and Eron decide to create a map to store their ingredients and their quantities. They start by declaring the map and its data type:

*/

var ingredients map[string]int

/*

Here, string represents the data type of the keys, which are the ingredient names, and int represents the data type of the values, which are the quantities of each ingredient.

*/

func main() {

/*

Alice and Eron then initialize the map by assigning it a value:

*/

ingredients = map[string]int{

"Verseium": 10,

"Sleseium": 20,

"Demiseium": 30,

}

/*

They then print the map to see the ingredients and their quantities:

*/

fmt.Println(ingredients)

}


/*

The output of the program is:

map[Verseium:10 Sleseium:20 Demiseium:30]

Alice and Eron are thrilled to see that their map is working as expected. They now know how to create a map and store key-value pairs in it. They can now use this map to keep track of the ingredients they have and the quantities they possess.

*/




