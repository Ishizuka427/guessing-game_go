package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10) + 1
	count := 0

	for {
		fmt.Print("Your guess? ")
		var guess int
		count++

		fmt.Scanf("%d", &guess)

		if guess == answer {
			fmt.Printf("Bingo! It took %v guesses!\n", count)
			return
		} else if guess < answer {
			fmt.Println("The Answer is higher!")
		} else {
			fmt.Println("The Answer is lower!")
		}
	}
}
