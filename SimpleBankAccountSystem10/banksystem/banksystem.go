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
	Deposit  int    `json:"deposit"`
}

func SaveUserAccountModification(UserName, Password string, Deposit int) {
	response := createUser{
		UserName: UserName,
		Password: Password,
		Deposit:  Deposit,
	}

	writeToJson, _ := json.Marshal(response)
	writeToJsonByte := []byte(writeToJson)
	err := os.WriteFile("userCredentials.json", writeToJsonByte, 0640)
	if err != nil {
		errors.New("Error")
	}
}

func CreateUser() {
	var UserName string
	var Password string
	var Deposit int
	fmt.Println("Enter user name:")
	fmt.Scan(&UserName)
	fmt.Println("Enter a password: ")
	fmt.Scan(&Password)
	fmt.Println("Enter a minimum 100$ deposit: ")
	fmt.Scan(&Deposit)
	if Deposit < 100 {
		fmt.Println("The amount entered is under the minimum allowed")
		fmt.Println("Please enter again")
		fmt.Scan(&Deposit)
	}

	SaveUserAccountModification(UserName, Password, Deposit)

}
func LogIn() (bool, string, string, int) {
	var UserName string
	var Password string
	fmt.Println("Enter user name:")
	fmt.Scan(&UserName)
	fmt.Println("Enter a password: ")
	fmt.Scan(&Password)

	credentialsValue, _ := os.ReadFile("userCredentials.json")
	var user createUser
	json.Unmarshal(credentialsValue, &user)

	if UserName == user.UserName && Password == user.Password {
		return true, UserName, Password, user.Deposit
	} else {

		return false, "", "", 0
	}
}

func LogOut() bool {
	return false
}
