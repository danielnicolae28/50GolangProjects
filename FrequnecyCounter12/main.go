package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Frequency counter")
	CreateText()
	CountWords()
}

func CreateText() {

	textTest := "This is a test so we can count each word we have on this string"

	file, err := os.Create("countWords.txt")
	if err != nil {
		panic("not a valid file")
	}

	writeToFile, err := file.WriteString(textTest)

	if err != nil {
		file.Close()

		panic("not a valid input")

	}
	fmt.Println(writeToFile)
}

func CountWords() {
	data, err := os.Open("./countWords.txt")

	if err != nil {
		panic("not a valid input")
	}
	defer data.Close()

	read := bufio.NewReader(data)

	line, _, err := read.ReadLine()

	getLine := []string{string(line)}

	for _, i := range getLine {
		fmt.Println(strings.Split(i, " "))
	}

	if len(line) > 0 {
		fmt.Println(string(line))
	}

	if err != nil {
		return
	}

}
