// This is to explain the concept of go mutexes and how they are used to manage concurrent operations

/*

In response to the attack on the kingdom, Alice and Eron realized that they needed to enhance their defense weapons with superpowers. They started by creating a Weapon struct that would hold the information for each weapon, including its name, damage, and superpower status:
*/

package main


import (

	"fmt"

	"sync"

)


type Weapon struct {
    Name       string
    Damage     int
    Superpower bool
}


/* 
Next, they created a slice of Weapon structs to hold the information for all of their defense weapons:
*/

var weapons []Weapon = []Weapon{
    {"Sword", 10, false},
    {"Bow", 8, false},
    {"Axe", 12, false},
    {"Shield", 5, false},
}


/*

To enhance the weapons with superpowers, Alice and Eron decided to use a goroutine to iterate through the slice of weapons and add a superpower to each weapon that met certain conditions.

*/

func addSuperpowers(weapons []Weapon, mutex *sync.Mutex) {
    mutex.Lock() // Acquire the mutex lock
    defer mutex.Unlock() // Release the mutex lock

    for i, weapon := range weapons {
        if weapon.Damage >= 10 {
            weapons[i].Superpower = true
        }
    }
}

func main() {
	var mutex sync.Mutex // Create a mutex lock

	// Add superpowers to weapons with damage >= 10
	go addSuperpowers(weapons, &mutex)

	// Wait for the goroutine to finish
	mutex.Lock()
	defer mutex.Unlock()

	// Print the weapons
	fmt.Println("Weapons:", weapons)
}


/*
To protect the shared data and resources in their program, Alice and Eron used a mutex to control access to the weapons slice. The addSuperpowers function is defined with a parameter of mutex that points to a sync.Mutex value.

Inside the function, the mutex lock is acquired using the Lock method, and the defer keyword is used to ensure that the mutex lock is released using the Unlock method after the function has finished executing.

The function then iterates through the weapons slice and adds a superpower to each weapon that has a damage value of 10 or higher.

To use the addSuperpowers function in their program, Alice and Eron created a new goroutine and passed in the weapons slice and the mutex value as arguments:

go addSuperpowers(weapons, &mutex)


By using a mutex to control access to the shared data and resources in their program, Alice and Eron were able to enhance their defense weapons with superpowers without causing conflicts or errors. The mutex ensured that only one goroutine could access the weapons slice at a time, preventing data races and ensuring the integrity of the data.

Through their experience with mutexes and the challenges they faced, Alice and Eron learned about the importance of properly managing access to shared data and resources in a programming project. Mutexes are a powerful tool that can be used to protect shared data structures from concurrent access, ensuring that the program runs smoothly and effectively.

*/