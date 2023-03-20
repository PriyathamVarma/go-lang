// This is to explain the concept of pointers in Go

/*

Alice and Eron have been working hard on building the kingdom's security system. They realized that they need to create a list of all the guards in the kingdom along with their names, ages, and ranks. To accomplish this, they create a struct that holds this information:

*/

package main

import "fmt"

type guard struct {
	name string
	age  int
	rank string
}

/*
Now they need to create a list of guards using this struct. They create a slice to hold the guard values:
*/

var guards = []guard{
	{"John", 30, "Captain"},
	{"Sarah", 28, "Lieutenant"},
	{"Mark", 35, "Captain"},
}

/*

Now, let's say Alice and Eron want to change the age of a guard in the list. They might try to do something like this:

guards[0].age = 33

This will work, but it's not the best way to do it. Instead, they should use a pointer to the guard struct. This will allow them to change the value of the guard struct in the list.

The reason is that the guards slice is actually holding copies of the guard structs, not the structs themselves. So when they try to change the age of a guard, they're only changing the age of the copy in the slice, not the original guard struct.

This is where pointers come in. A pointer is a variable that stores the memory address of a value, instead of the value itself. In Go, the * symbol is used to denote a pointer.

So instead of storing copies of the guard structs in the guards slice, Alice and Eron can store pointers to the guard structs:

*/

var guardsWithPointers = []*guard{
	&guard{"John", 30, "Captain"},	// Notice the & symbol
	&guard{"Sarah", 28, "Lieutenant"}, 
	&guard{"Mark", 35, "Captain"},
}

/*

Now, when they try to change the age of a guard, they can use the pointer to the guard struct:

*/

func main() {

	guardsWithPointers[0].age = 33
	fmt.Println(guardsWithPointers[0].age)

}

/*

In this example, Alice and Eron are using the & symbol to get the memory address of a guard struct.
This is called the address operator. The * symbol is used to get the value of a pointer. This is called the dereference operator.


/*

