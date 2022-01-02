package main

import (
	"fmt"
)

func main() {
	greeter()

	result := add(10, -19)
	fmt.Println("Result 0 == ", result)

	result1 := proAdder(9)
	fmt.Println("Result 1 == ", result1)

	result2 := proAdder()
	fmt.Println("Result 2 == ", result2)

	result3 := proAdder(-10, 1, 0, 7)
	fmt.Println("Result 3 == ", result3)

	resultStr, result4 := multipleReturns(15)
	fmt.Printf("Result 4 == %v and Result String is - %v\n", result4, resultStr)

}

func greeter() {
	fmt.Println("Welcome to functions in Golang!")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func multipleReturns(number int) (string, int) {
	return "Returning from multipleReturns function", number * number
}
