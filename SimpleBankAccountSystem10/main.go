package main

import (
	"fmt"

	b "github.com/danielnicolae28/50GolangProjects/SimpleBankAccountSystem10/banksystem"
)

/*
use structs, interfaces and maps, slices
*/
func main() {
	var userChoice int
	fmt.Println("Simple bank account system")

	fmt.Println("Welcome to Go Bank")
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:

		b.CreateUser()

	case 2:

		b.LogIn()

	}

}
