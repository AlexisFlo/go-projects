package main

import "fmt"

/*

	1. What is the result of the following program? Why?


func main() {
	var userName
	userName = "user"
	fmt.Println(userName)
}

syntax error: unexpected newline, expected type
Error. Go is statically typed and you need to tell Go Compiler the data type when doing variable declaration
*/

/*
	2. What is the result of the following program? Why?

func main() {
	var userName = 2
	fmt.Println(userName)
}

It will print 2. The reason it doesn't fail (even though you didn't declare the type) is due to Go performing type inference where infers the type from the assigned value

*/

/*
3. Fix the following program by modifying one of the lines (but not adding or removing lines)

	func main() {
		var userName
		userName = "user"
		fmt.Println(userName)
	}
*/
// func main() {
// 	var userName string // declare the type of userName
// 	userName = "user"
// 	fmt.Println(userName)
// }

/*
	4. Modiy the following program to print the types of the variables

	func main() {
		var food = "Pizza"
		var slices = 4
		var pineappleOnPizza = True
	}
*/

func main() {
	var food = "Pizza"
	var slices = 4
	var pineappleOnPizza = true

	fmt.Printf("food is %T\nslices is %T\npinneapleOnPizza is %T", food, slices, pineappleOnPizza)
}
