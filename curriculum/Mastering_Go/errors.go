// This is to explain the concept of go errors

/*

Alice and Eron were now masters of Go programming, but they knew that even the best programmers make mistakes. They also knew that errors were a natural part of any programming project, and they needed to know how to handle them effectively.

One day, while they were working on a new project to build a system that could predict the weather, they encountered an error. The program was crashing whenever it encountered an invalid input. They realized that they needed to handle errors in a more effective way.

They started by creating an error type called InvalidInputError that would be returned whenever the program encountered an invalid input:

*/


type InvalidInputError struct {
    input string
    msg   string
}


