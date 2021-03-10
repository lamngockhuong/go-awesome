package main

import "fmt"

func main() {
	fmt.Print("Enter first number: ")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Print("Enter second number: ")
	var secondNumber int
	fmt.Scan(&secondNumber)

	firstNumber = firstNumber - secondNumber
	secondNumber = firstNumber + secondNumber
	firstNumber = secondNumber - firstNumber

	println("First number: ", firstNumber)
	println("Second number: ", secondNumber)
}
