package main

import "fmt"

func main() {
	numbersSlice := []int{10, 20, 30}
	stringsSlice := []string{"Hello,", "\n"}

	fmt.Println(len(numbersSlice))
	fmt.Println(len(stringsSlice))
	numbersSlice = append(numbersSlice, 40)
	numbersSlice = append(numbersSlice, 50)
	stringsSlice = append(stringsSlice, "How")
	stringsSlice = append(stringsSlice, "are")
	stringsSlice = append(stringsSlice, "you?")

	fmt.Println(sumSlice(numbersSlice))
	fmt.Println(printSliceInOneLine(stringsSlice))

	numbersSlice2 := []int{20, 70, 90}
	numbersSlice = append(numbersSlice, numbersSlice2...)
	fmt.Println(sumSlice(numbersSlice))

}

func sumSlice(slice []int) int {
	sum := 0
	for _, number := range slice {
		sum += number
	}
	return sum
}

func printSliceInOneLine(slice []string) string {
	text := ""
	for i, word := range slice {
		if slice[i] != "\n" {
			text += word + " "
		}
	}
	return text
}
