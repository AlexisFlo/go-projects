package main

import "fmt"

/*
1. Define a function that gets two integers, adds them and returns the result
	i. Make sure to also write code that executes the function
*/

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Result:", add(2, 12))
}

/*
2. What is the problem with the following code? How to fix it?

func add(x, y int) {
	return x + y
}

func main() {
	fmt.Println("Result:", add(2, 3))
}

It returns an integer but at the same time the function doen't specify any return value, so Go expects the function to return nothing
*/