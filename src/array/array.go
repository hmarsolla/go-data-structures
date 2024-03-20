package main

import "fmt"

func main() {

	numbersArray := [...]int{10, 20, 30, 40, 50}
	stringsArrays := [...]string{"Hello,", "\n", "this", "is", "my", "message", "to", "you!", "\n", "How are you?"}

	fmt.Println(len(numbersArray))
	fmt.Println(len(stringsArrays))
	fmt.Println(sumArray(numbersArray))
	fmt.Println(printArrayInOneLine(stringsArrays))
}

func sumArray(array [5]int) int {
	sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

func printArrayInOneLine(array [10]string) string {
	text := ""
	for i, word := range array {
		if array[i] != "\n" {
			text += word + " "
		}
	}
	return text
}
