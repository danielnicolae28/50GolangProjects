package managetodos

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var userTasksStorage []string

func AddTodo() {
	var userInputTask string
	fmt.Println("enter your task: ")
	scann := bufio.NewScanner(os.Stdin)

	if scann.Scan() {
		userInputTask = scann.Text()
		userTasksStorage = append(userTasksStorage, userInputTask)
	}
}

func DeleteTodo() {
	var userDeleteTask string
	fmt.Println("enter the task you want to delete")
	scann := bufio.NewScanner(os.Stdin)
	if scann.Scan() {
		userDeleteTask = scann.Text()
	}
	userTasksStorage = slices.DeleteFunc(userTasksStorage, func(userTask string) bool {
		return userTask == userDeleteTask
	})

}

func DisplayTodo() {

	if len(userTasksStorage) == 0 {
		fmt.Println("You don't have any task")
	} else {

		fmt.Println(userTasksStorage)
	}

}
