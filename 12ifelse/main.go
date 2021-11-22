package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Plz enter age - ")
	var age int
	fmt.Scanln(&age)

	var result string

	if age >= 18 {
		result = "Permanent Lisence."
	} else if age >= 16 && age < 18 {
		result = "Learners Permit Only."
	} else {
		result = "Not eligible for driving."
	}

	fmt.Printf("Age == %v. %v\n", age, result)

	if age%2 == 0 {
		fmt.Println("Age is Even!")
	} else {
		fmt.Println("Age is Odd!")
	}

	fmt.Println("Plz enter some number - ")

	var someString string
	fmt.Scanln(&someString)

	if newNum, err := strconv.Atoi(someString); err == nil {
		fmt.Printf("New Number == %d, type: %T\n", newNum, newNum)
	}

}
