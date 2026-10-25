package main

/*
Generate a number between 0 and 10 and using switch statement do the following:
1. If it's 1 print "One is the loneliest number that you'll ever do"
2. If it's 2 print "Two can be as bad as one. It's the loneliest number since the number one"
3. For any other number print "Sorry, no mention of the number in the song"
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch randNum := rng.Intn(10); randNum {
	case 1:
		fmt.Println("One is the loneliest number that you'll ever do")
	case 2:
		fmt.Println("Two can be as bad as one. It's the loneliest number since the number one")
	default:
		fmt.Printf("Sorry, no mention of the number in the song: %v", randNum)
	}

}
