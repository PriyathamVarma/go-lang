# Basic GO quests

## Quest 1: Initial Setup

- Step 1: Install Go

Install Go on your local machine. You can download the latest version of Go from the [official website](https://golang.org/dl/).

- Step 2: Set up your Go workspace

Set up your Go workspace. You can follow the instructions on the [official website](https://golang.org/doc/code.html#Workspaces).

- Step 3: Set up your Go path

Set up your Go path. You can follow the instructions on the [official website](https://golang.org/doc/code.html#GOPATH).

- Step 4: Set up your Go editor

Set up your Go editor. You can follow the instructions on the [official website](https://golang.org/doc/code.html#Text_editor).

- Step 5: Set up your Go environment

Set up your Go environment. You can follow the instructions on the [official website](https://golang.org/doc/code.html#Environment).

- Step 6: Set up your Go version control

Set up your Go version control. You can follow the instructions on the [official website](https://golang.org/doc/code.html#Version_control).

**Sample**

```bash

>> Open code editor

>> Create a file named hello.go

>> Write the following code in the file

package main

import "fmt"

func main() {

    fmt.Println("Hello, World!")

}

>> Save the file

>> Open terminal

>> Navigate to the directory where the file is saved

>> Run the following command

go run hello.go

```

## Go Packages

In Go programming, a package is like a box that contains useful tools that we can use to make our programs do cool things.

Just like you have a toolbox at home with different tools like a hammer or screwdriver, Go packages have different tools that we can use to write better code.

For example, one package is called "fmt", which helps us format and print text on the screen. Another package is called "math", which helps us do math calculations like adding or subtracting numbers.

By using these packages, we don't have to write all the code ourselves from scratch. We can use the tools that the packages provide, and save time and effort.

And just like how you can borrow tools from your neighbor if you need something specific, Go programmers can borrow packages from other programmers and use them in their own code.

So, in short, Go packages are like toolboxes full of tools that help us write better and more efficient code!

## What is package main?

In Go programming, the package main is a special package that is used to create a standalone executable program.

When we create a Go program, we put all our code into one or more packages. The package keyword is used to declare the name of the package at the beginning of each source file.

However, when we create a standalone executable program, we need to have a special package called package main. This package tells Go that this is a standalone program that can be run on its own.

The package main must include a special function called main(). This function serves as the entry point of the program, which means that it's the first function that gets called when the program runs.

Inside the main() function, we can put all the code that we want to execute when the program runs. This could include things like printing messages to the console, performing calculations, or interacting with the user.

So, in summary, the package main is a special package in Go that is used to create standalone executable programs. It includes a special main() function that serves as the entry point of the program.
