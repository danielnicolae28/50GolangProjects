package banksystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type createUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func UserManagement(userChoice int) {

	switch userChoice {
	case 1:
		CreateUser()
	case 2:

		LogIn()

	}
}

func CreateUser() {
	var UserName string
	var Password string
	fmt.Println("Enter user name:")
	fmt.Scan(&UserName)
	fmt.Println("Enter a password: ")
	fmt.Scan(&Password)

	response := createUser{
		UserName: UserName,
		Password: Password,
	}

	writeToJson, _ := json.Marshal(response)
	writeToJsonByte := []byte(writeToJson)
	err := os.WriteFile("userCredentials.json", writeToJsonByte, 0640)
	if err != nil {
		errors.New("Error")
	}
	fmt.Println(writeToJson)

}
func LogIn() bool {

}

func LogOut() bool {

}
