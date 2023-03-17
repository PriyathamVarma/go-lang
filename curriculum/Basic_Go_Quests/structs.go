// This is to explain the concept of structs in Go

/*

After successfully analyzing the mysterious potion, Alice and Eron realized that they needed a better way to organize and store information about their potions. They needed to keep track of various properties of each potion, such as the name, the ingredients used, and the effects of the potion.

To solve this problem, Alice and Eron decided to use structs in Go programming. A struct is a user-defined type that groups together zero or more values of different types into a single entity.

They started by declaring a struct type called Potion, which contains fields for the name, ingredients, and effects of a potion:

*/

package main

import (
"fmt"
)

type Potion struct {
Name string
Ingredients []string
Effects []string
}


/*

Next, they created a new Potion using the struct type, filling in the fields with the appropriate values:

*/

func main() {

// var potion1 Potion

potion1 := Potion{
Name: "Strength Potion",
Ingredients: []string{"giant's toe", "troll's blood", "dragon scales"},
Effects: []string{"increased strength", "improved endurance"},
}

/*

OR

var potion1 Potion

potion1.Name = "Strength Potion"

potion1.Ingredients = []string{"giant's toe", "troll's blood", "dragon scales"}

potion1.Effects = []string{"increased strength", "improved endurance"}

*/

/*

They repeated the process to create additional potions, each with their own unique name, ingredients, and effects.

Now, with their potions stored in structs, Alice and Eron had a better way to organize and manage information about their creations. They could easily access and modify the properties of each potion using dot notation:

*/

fmt.Println(potion1.Name)

fmt.Println(potion1.Ingredients)

fmt.Println(potion1.Effects)

}


/*

Through this example, Alice and Eron learned the concept of structs in Go programming, 
which allowed them to group related data together into a single entity. With structs, 
they were able to better organize and manage information about their potions.


*/

