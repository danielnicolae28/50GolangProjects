package main

import (
	"fmt"

	a "github.com/danielnicolae28/50GolangProjects/SimpleBankAccountSystem10/accountSettings"
	b "github.com/danielnicolae28/50GolangProjects/SimpleBankAccountSystem10/banksystem"
)

/*
use structs, interfaces and maps, slices
*/
func main() {
	var userChoice int
	fmt.Println("Simple bank account system")

	fmt.Println("Welcome to Go Bank")
	fmt.Println("1 for create user")
	fmt.Println("2 log in")
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		b.CreateUser()
	case 2:

		isLogIn, userName, password, Deposit := b.LogIn()
		if isLogIn {
			a.AccountOn(userName, password, Deposit)
		} else {
			isLogIn = b.LogOut()
		}

	}

}
