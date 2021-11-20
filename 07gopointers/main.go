package main

import "fmt"

func main() {
	var onePtr *int
	fmt.Println("Value of onePtr == ", onePtr) // <nil>

	myNumber := 15

	var ptr = &myNumber
	fmt.Println("Value of actual pointer == ", ptr)  // random Location == 0xc0000b2008
	fmt.Println("Value of actual pointer == ", *ptr) // 15

	*ptr = *ptr + 5
	fmt.Println("Updated Value == ", myNumber) // 20

}
