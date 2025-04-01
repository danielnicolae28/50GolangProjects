package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Palindrome Checker")
	var storeResult []string
	var reverse []string
	var userInput string
	fmt.Println("Enter the palindrome word: ")
	fmt.Scan(&userInput)

	for _, val := range userInput {
		storeResult = append(storeResult, string(val))
	}
	for i, _ := range storeResult {
		reverse = append(reverse, storeResult[len(storeResult)-1-i])
	}

	if reflect.DeepEqual(storeResult, reverse) {
		fmt.Println("Is palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}

}
