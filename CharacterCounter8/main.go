package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var userInputString string
	var userInputSearch string

	fmt.Println("Character Counter")
	scann := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string:")

	if scann.Scan() {
		userInputString = scann.Text()
	}

	fmt.Println("Enter which character to search:")
	fmt.Scan(&userInputSearch)

	// fmt.Println(s.Count(userInputSearch, userInputSearch))

	var storeData []string

	for _, val := range userInputString {
		storeData = append(storeData, string(val))
	}
	i := 0
	for _, data := range storeData {
		if data == userInputSearch {
			i++
		}

	}

	fmt.Printf("The word you searched was find %v times\n", i)
}
