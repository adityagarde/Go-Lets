package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to maths in Go!")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1) // Random number between [1 and 5]

	// crypto/rand
	//myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	//fmt.Println(myRandomNum)
}
