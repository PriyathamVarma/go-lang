// This is to explain the concept of go channels and how they are used to manage concurrent operations

/*

Alice and Eron had successfully implemented the use of goroutines to monitor multiple areas of the castle, 
making it much more difficult for intruders to breach the security system. However, one day, the castle came 
under attack from a group of ruthless attackers who were determined to break into the kingdom and steal the valuable Verseium metal.

The attackers were using a sophisticated hacking technique that targeted the security system's communication channels, 
causing them to overload and break down. Alice and Eron knew they had to act fast to come up with a solution.

They realized that they needed a way to communicate between the goroutines that were monitoring different areas 
of the castle to coordinate their efforts and respond to the attackers' attack. That's when they learned about Go channels, 
a powerful feature of Go programming that allows for communication and synchronization between goroutines.

To demonstrate the concept of Go channels, Alice and Eron decided to use an example of a factory that produced different products. 
Each product had to go through multiple stages of production before it was completed, and each stage had to communicate with the next to 
ensure that the final product was produced correctly.

They started by creating a function called produceProduct that would simulate the production of a product and send it through a channel to the next stage:


*/

func produceProduct(product string, stage int, out chan<- string) {
    // Perform the production process for the product and send it to the next stage
    out <- fmt.Sprintf("%s is produced in stage %d", product, stage)

	fmt.
}


/*

Next, they needed to simulate the production of multiple products. To do this, they created a slice of products and used a loop to produce each product:

*/


func main(){

	// proudtcs
	var products []string = []string{"book", "computer", "chair", "table", "lamp"}



	// Channels
	var stage1 chan string = make(chan string)
	stage2 := make(chan string)
	stage3 := make(chan string)
	stage4 := make(chan string)


	// for loop and passing the channel to the function

	for _, product := range products {
		go produceProduct(product, 1, stage1)
	}

	for _, product := range products {
		go produceProduct(product, 2, stage2)
	}

	for _, product := range products {

		go produceProduct(product, 3, stage3)
	}

	for _, product := range products {

		go produceProduct(product, 4, stage4)
	}


}

/*

With this change, each function call would run concurrently and communicate through channels, allowing for multiple products to be produced at the same time and ensuring that each stage received the correct input.

As Alice and Eron continued to experiment with Go channels, they realized that they could apply this concept to their security system. They could use channels to coordinate the efforts of the different goroutines that were monitoring different areas of the castle and ensure that they were communicating effectively to respond to the attackers' attack.

With their newfound knowledge of Go channels and the power of concurrency, Alice and Eron worked together to quickly implement a new security system that used channels to coordinate the efforts of the different goroutines. They were able to respond to the attackers' attack and prevent them from breaching the security system.

Through their experience with Go channels and the challenges they faced, Alice and Eron learned about the importance of effective communication and synchronization between goroutines in a programming project. They also learned the importance of being able to quickly adapt to new challenges and come up with innovative solutions to protect the kingdom and its valuable Verseium metal.

*/