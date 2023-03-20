// This is to explain the concept of goroutines in Go

/*
Alice and Eron were now the lead programmers for the kingdom's security system. They had used their knowledge of Go programming to build a strong foundation, but they knew they needed to take their skills to the next level.

One day, while they were working on the system, the castle's guard captain approached them with a new problem. The kingdom had been experiencing a surge in security breaches, and the guards needed to be able to monitor multiple areas of the castle at the same time.

Alice and Eron realized that they needed a way to perform multiple tasks concurrently to solve this problem. They soon learned about goroutines, a powerful feature of Go programming that allows for concurrent execution of functions.

To demonstrate the concept of goroutines, Alice and Eron decided to use an example of managing a warehouse. They had multiple products that needed to be moved from one location to another, and they needed to do this as efficiently as possible.

They started by creating a function called moveProduct that would simulate the movement of a product from one location to another:
*/

package main

import (
	"fmt"

	"time"
)

func moveProduct(product string, fromLocation string, toLocation string) {
	fmt.Printf("Moving product %s from %s to %s\n", product, fromLocation, toLocation)
}


/*

Next, they needed to simulate moving multiple products. To do this, they created a slice of products and used a loop to move each product:

*/

func main(){
	var products []string = []string{"book", "computer", "chair", "table", "lamp"}

	// This code can be ignored
	for _, product := range products {
		moveProduct(product, "Shelf 1", "Shelf 2")
	}

	/*
	This code would move each product from "Shelf 1" to "Shelf 2" one at a time. However, Alice and Eron knew that they could use goroutines to move multiple products at the same time, making the process more efficient.

	They modified the code to use goroutines by adding the go keyword in front of the function call:
	*/


	for _, product := range products {
		go moveProduct(product, "Shelf 1", "Shelf 2")
	}


}

/*

With this change, each function call would run concurrently, allowing for multiple products to be moved at the same time.

As Alice and Eron continued to experiment with goroutines, they realized that they could apply this concept to the castle's security system. They could use goroutines to simultaneously monitor multiple areas of the castle, making it much more difficult for intruders to breach the security system.

However, they soon encountered a plot twist. As they implemented the new system using goroutines, 
they discovered that the concurrent execution of functions could cause data races and other issues if not managed properly. 
They needed to use additional techniques such as mutexes and channels to ensure that their system was secure and reliable.

Through their experience with goroutines and the challenges they faced, Alice and Eron learned about the importance of concurrency
 and the need for proper management of concurrent operations in their programming projects.

*/




