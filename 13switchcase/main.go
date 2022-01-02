package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Golang!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	// fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 place.")
	case 2:
		fmt.Println("You can move 2 places.")
	case 3:
		fmt.Println("You can move 3 places.")
		fallthrough
	case 4:
		fmt.Println("You can move 4 places.")
		fallthrough
	case 5:
		fmt.Println("You can move 5 places.")
	case 6:
		fmt.Println("You can move 6 places and roll the dice again!")
	default:
		fmt.Println("Unreachable code!")
	}
}
