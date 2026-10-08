/*
Conditionals
are used to perform different actions based on different conditions. In Go, the primary conditional statements are 'if', 'else', and 'switch'.

1. Generate a random number between 1 and 100
	i. If the number is higher than 50 print "It's closer to 100"
	ii. If the number is lower than 50 print "It's closer to 0"
	iii. Print the generated random number
2. Modify the previous code to print "It's 50!" if the random number is 50
3. Modify the conditional in the code you previously written to check not only if the number is higher than 50 but also if's it's even, If it's even and higher than 50, print "It's closer to 100 and it's even"
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var randNum int = rng.Intn(100)

	if randNum > 50 {
		fmt.Println("It's closer to 100")
	} else {
		fmt.Println("It's closer to 0")
	}
	fmt.Printf("Generated number: %v\n", randNum)
}
