package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go!")

	var fruitList [4]string

	fruitList[0] = "Apples"
	fruitList[1] = "Mangoes"
	fruitList[3] = "Bananas"

	fmt.Println("fruitList is == ", fruitList)          // [Apples Mangoes  Bananas] - Space for 2nd location
	fmt.Println("fruitList length == ", len(fruitList)) // 4

	var vegList = [3]string{"potato", "beans", "carrot"}
	fmt.Println("Vegetable list == ", vegList)
	fmt.Println("vegList length == ", len(vegList)) // 3

}
