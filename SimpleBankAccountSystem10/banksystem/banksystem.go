package banksystem

import "fmt"

type createUser struct {
	UserName string
	Password string
}

func (c createUser) CreateUser() {
	fmt.Println("Enter user name:")
	fmt.Scan(&c.UserName)
	fmt.Println("Enter a password: ")
	fmt.Scan(&c.Password)

}
func LogIn() {

}

func LogOut() {

}
