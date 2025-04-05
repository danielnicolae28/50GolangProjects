package main

import (
	"fmt"

	m "github.com/danielnicolae28/50GolangProjects/SimpleTodoList9/managetodos"
)

func main() {
	fmt.Println("Simple Todo List")
	var userChoice int
	for {

		fmt.Println("1. Add tsk")
		fmt.Println("2. Delete tsk")
		fmt.Println("3. Display List")
		fmt.Println("Press any key to exit the list")

		fmt.Println("\n Choice:")
		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			m.AddTodo()
		case 2:
			m.DeleteTodo()
		case 3:
			m.DisplayTodo()
		default:
			fmt.Println("Not a valid input, please enter a number from the choice displayed")
			break
		}
	}

}
