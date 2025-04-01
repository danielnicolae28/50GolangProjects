package main

import "fmt"

func main() {
	fmt.Println("FizzBuzz")

	for i := 0; i <= 100; i++ {
		if (i % 3) == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}
