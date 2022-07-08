package main

import (
	"fmt"
	"strconv"
)

//take two numbers and return the sum of them
func sum(a int, b int) int {
	return a + b
}

//take two numbers and return the subtraction of them
func sub(a int, b int) int {
	return a - b
}

//take two numbers and return the multiply of them
func mul(a int, b int) int {
	return a * b
}

//take two numbers and return the division of them
func div(a int, b int) int {
	return a / b
}

//change string to int
func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

//change int to string
func intToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

//take in two integers and a math operator and return the result
func math(a int, b int, operator string) int {
	switch operator {
	case "+":
		return sum(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return mul(a, b)
	case "/":
		return div(a, b)
	default:
		return 0
	}
}

func readLine() string {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		panic(err)
	}
	return str
}

func main() {
	println("Enter two numbers:")
	a := stringToInt(readLine())
	b := stringToInt(readLine())
	println("Enter an operator:")
	operator := readLine()
	println("Result: " + intToString(math(a, b, operator)))
}
