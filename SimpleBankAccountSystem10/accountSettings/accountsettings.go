package accountsettings

import (
	"fmt"

	b "github.com/danielnicolae28/50GolangProjects/SimpleBankAccountSystem10/banksystem"
)

func AccountOn(userName, password string, depositAmount int) {
	var userChoose int
	fmt.Printf("Welcome %v\n", userName)
	fmt.Printf("Your account has %v$\n", depositAmount)
	fmt.Println("Choose an option:")
	fmt.Println("1.Withdraw")
	fmt.Println("2.Deposit")
	fmt.Println("3.Exit")
	fmt.Scan(&userChoose)

	switch userChoose {
	case 1:
		Withdraw(userName, password, depositAmount)
	case 2:
		Deposit(userName, password, depositAmount)
	default:
		break

	}
}

func Withdraw(userName, password string, depositAmount int) {
	var userInput int

	fmt.Println("Enter an amount")
	fmt.Scan(&userInput)

	if userInput > depositAmount {
		fmt.Println("Not a valid amount, plese enter again")
		Withdraw(userName, password, depositAmount)

	} else {
		depositAmount -= userInput
		b.SaveUserAccountModification(userName, password, depositAmount)

		fmt.Printf("the current amount is %v", depositAmount)
	}
}
func Deposit(userName, password string, depositAmount int) {
	var userInput int
	fmt.Println("Enter deposit amount")
	fmt.Scan(&userInput)

	if userInput <= 0 {
		fmt.Println("plese enter a positiv value")
		Deposit(userName, password, depositAmount)
	} else {
		depositAmount += userInput
		b.SaveUserAccountModification(userName, password, depositAmount)
		fmt.Printf("the current amount is %v", depositAmount)
	}

}
