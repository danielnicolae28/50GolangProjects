package main

import "fmt"

func main() {
	var userInput int
	fmt.Println("Multiplication table ")
	fmt.Println("Insert a number for multiplication table:")
	fmt.Scan(&userInput)

	var multiplicationTabel = [2][3]int{{1, 2, 3}, {5, 6, 7}}

	for _, v := range multiplicationTabel {
		for _, j := range v {
			j = j * userInput
			fmt.Println(j)
		}
	}

	fmt.Println(multiplicationTabel)
}
