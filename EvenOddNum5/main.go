package main

import "fmt"

func main() {
	fmt.Println("EvenOdd num")
	var userInput int

	fmt.Println("Enter a number")
	fmt.Scan(&userInput)

	if (userInput % 2) == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
