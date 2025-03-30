package main

import "fmt"

func main() {

	var temp float64
	var response string

	fmt.Println("Temperature Convertor")

	fmt.Println("If you want to convert Celsius to Fahrenheit press F otherwise press C")
	fmt.Scan(&response)

	if response == "F" || response == "f" {
		fmt.Println("Enter Celsius temperature to be coverted in Fahrenheit")
		fmt.Scan(&temp)
		f := (temp * 9 / 5) + 32

		fmt.Println(f)

	} else if response == "C" || response == "c" {
		fmt.Println("Enter Fahrenheit temperature to be converted in Celsius")
		fmt.Scan(&temp)
		c := (temp - 32) / 1.8
		fmt.Println(c)
	}
}
