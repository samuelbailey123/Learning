package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}

	digits = append([]int{1}, digits...)

	return digits
}

func main() {
	// fmt.Println(plusOne([]int{1, 2, 3}))
	myArray := []int{9}
	myArray = append([]int{1}, myArray...)
	myArray[1] = 0
	fmt.Println(myArray)
}
