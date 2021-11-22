package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go!")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Println("fruitList == ", fruitList)
	fmt.Printf("Type of fruitList == %T\n", fruitList) // []string i.e. slice

	fruitList = append(fruitList, "Strawberry", "Watermelon", "Muskmelon")
	fmt.Println("Updated fruitlist == ", fruitList) // [Apple Mango Banana Strawberry Watermelon Muskmelon]

	fruitList = append(fruitList[1:])
	fmt.Println("New fruitList == ", fruitList) // [Mango Banana Strawberry Watermelon Muskmelon] - Start from 1 till end

	fruitList = append(fruitList[1:4])
	fmt.Println("New fruitList == ", fruitList) // [Banana Strawberry Watermelon] - Start from 1 to 4, 4 not included

	fruitList = append(fruitList[:2], "Peach")
	fmt.Println("New fruitList == ", fruitList) // [Banana Strawberry Peach] - Starts from 0 to 1, All elements before 2

	highScores := make([]int, 4) // (Datatype, Size)

	highScores[0] = 434
	highScores[1] = 933
	highScores[2] = 234
	highScores[3] = 882
	// highScores[4] = 899 // runtime error: index out of range [4] with length 4

	highScores = append(highScores, 739, 110, 98, 777) // accepts these values without error and appends them

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // false
	sort.Ints(highScores)
	fmt.Println(highScores)                     // [98 110 234 434 739 777 882 933]
	fmt.Println(sort.IntsAreSorted(highScores)) // true

	// Removing values from Slices using index

	var courses = []string{"SpringBoot", "Quarkus", "React.js", "Micronaut", "Helidon", "Vert.x"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...) // append is used as it re-allocated memory
	fmt.Println(courses)                                    // [0-index) && [index+1, end]
}
