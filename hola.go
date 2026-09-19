package main

import "fmt"

func main() {
	// const name = "Alexis"
	// age := 26
	// fmt.Println("My name is", name, "and I'm", age, "years old")

	// fmt.Println("Skynet Beta Testing")

	// var Objective = "Defend Humanity"
	// This fail because the objective vairable is not used

	const squares = 2
	var circles = 0

	fmt.Println("Squares:", squares)
	fmt.Println("Circles:", circles)

	squares = 1
	cirlces = 7

	fmt.Println("Squares:", squares)
	fmt.Println("Circles:", circles)

	// this will fail because the variable squares is a constant and
	// cannot be reassigned a new value

}
