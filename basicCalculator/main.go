package main

import "fmt"

func main() {

	var num1, num2, result float64
	var operation string
	var storeResult []float64

	fmt.Println("Basic Calculator")
	for {

		fmt.Println("Enter first number :")
		fmt.Scan(&num1)
		fmt.Println("Enter second number :")
		fmt.Scan(&num2)

		fmt.Println("Choose the operation (+,-,*,/) : ")
		fmt.Scan(&operation)

		switch operation {
		case "+":
			result = num1 + num2
			storeResult = append(storeResult, result)
			fmt.Printf("Sum result : %v\n", result)
		case "-":
			result = num1 - num2
			storeResult = append(storeResult, result)
			fmt.Printf("Substraction result : %v\n", result)
		case "*":
			result = num1 * num2
			storeResult = append(storeResult, result)
			fmt.Printf("Multiplication result : %v\n", result)
		case "/":
			result = num1 / num2
			storeResult = append(storeResult, result)
			fmt.Printf("Devision result : %v\n", result)
		}

		fmt.Printf("Result history, %v\n", storeResult)

		var response string
		fmt.Println("Would you want another calculation : yes or no")
		fmt.Scan(&response)
		if response == "yes" || response == "YES" || response == "Yes" {
			continue
		} else {
			break
		}

	}

}
