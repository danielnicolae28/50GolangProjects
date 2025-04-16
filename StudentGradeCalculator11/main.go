package main

import "fmt"

func main() {
	fmt.Println("Student grade calculator")
	GradesInput()
}

func GradesInput() {
	var gradesNum int
	var gradesStorage []float64
	var grade float64
	var sum float64

	count := 0
	fmt.Println("how many grades do you have?")
	fmt.Scan(&gradesNum)

	for i := 1; i <= gradesNum; i++ {

		fmt.Printf("enter grade number %v:  ", i)
		fmt.Scan(&grade)
		gradesStorage = append(gradesStorage, grade)
	}

	for _, v := range gradesStorage {
		count++
		sum = sum + v

	}
	average := sum / float64(count)
	fmt.Printf("The verage of your grades is : %.2f\n", average)

}
