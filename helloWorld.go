package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)
	fmt.Println("Input your guess")
	fmt.Println(number)
	var guess int = 1

	for {
		fmt.Println("Input your guess")
		fmt.Scan(&guess)
		if guess > number {
			fmt.Println("Your guess is bigger")
		} else if guess < number {
			fmt.Println("Your guess is smaller")
		} else {
			fmt.Println("You win")
			break
		}
	}

}
