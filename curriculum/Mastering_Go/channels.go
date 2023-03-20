// This is to explain the concept of go channels and how they are used to communicate between goroutines

package main

import (
	"fmt"
	"time"

	"sync"
)

func main() {

	/*
		Channels are a way for goroutines to communicate with each other. They are a typed conduit through which you can send and receive values with the channel operator, <-.

		Here's an example of a channel of strings buffering up to 2 values.
	*/

	// This is a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	/*
		We can send a value into a channel using the channel <- syntax. Here we send "buffered" and "channel" into the messages channel from different goroutines.
	*/

	// We can send a value into a channel using the channel <- syntax.
	go func() { messages <- "buffered" }()
	go func() { messages <- "channel" }()

	/*
		We can receive a value from a channel using the <-channel syntax. Here we receive the values "buffered" and "channel" as we previously sent them on the channel.

		Note that the order of the received values corresponds to the order of the sends.

		We can also use the range keyword to iterate over values received from a channel. When we close a channel, the iteration terminates after receiving the last sent value. If the channel is not closed, we'll receive the zero value for the channel's type.

		Here we iterate over 2 values in the messages channel and then close it.

	*/



}







