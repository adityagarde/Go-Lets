package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	defer fmt.Println("World!")

	defer fmt.Println("One")
	fmt.Println("Two")
	defer fmt.Println("Three")
	defer fmt.Println("Four")

	defer fmt.Println("Hello!")
	fmt.Println("World!")

	godefer()

	/*
		Defer statements executed in LIFO fashion
			Hello!
			Two
			World!
			43210Hello!
			Four
			Three
			One
			World!
	*/
}

func godefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
