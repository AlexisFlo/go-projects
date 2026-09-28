package main

/*
Packages

Objectives
	1. Using the time package, print The time now is: <TIME>
	2. Using the math/rand package, generate random integer between 0 and 100
	3. Using the math package calculate the square of 9
*/
import (
	"fmt"
	"time"
	"math"
	"math/rand"
)


func main() {

	fmt.Println("The time now is:", time.Now())
	fmt.Println("Random number", rand.Intn(100))
	fmt.Printf("The square of 9: %g\n", math.Sqrt(9))
}
