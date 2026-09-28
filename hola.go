package main

import "fmt"

/*
1. Modify the following program to ask for a name and an age
	i. Modify the print statement accordingly

func main() {
	fmt.Printf("Your name is __ and your age is __")
}
*/

func main() {
	var firstName string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	
	fmt.Printf("Your name is %v and your age is %v", firstName, age)
}
