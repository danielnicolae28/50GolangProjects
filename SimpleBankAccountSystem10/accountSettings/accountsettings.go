package accountsettings

import (
	"fmt"
)

func AccountOn(userName string, depositAmount int) {
	var userChoose int
	fmt.Printf("Welcome %v\n", userName)
	fmt.Printf("Your account has %v$\n", depositAmount)
	fmt.Println("Choose an option:")
	fmt.Scan(&userChoose)

	switch userChoose {
	case 1:
		Withdraw()
	case 2:
		Deposit()
	case 3:
		break
	default:
		break

	}
}

func Withdraw() {

}
func Deposit() {

}
