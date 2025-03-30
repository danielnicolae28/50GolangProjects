package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Guessing Game")
	randNum := rand.Intn(10)
	var num int
	fmt.Println(randNum)
	for {

		fmt.Print("choose a number between 1 and 10:  ")
		fmt.Scan(&num)

		if randNum > num {
			fmt.Println("Your num is to small, try with another guess")
		} else if randNum < num {
			fmt.Println("Your num is to big, try with another guess")
		} else {
			fmt.Println("You guessed, nice!")
			break
		}
	}
}
