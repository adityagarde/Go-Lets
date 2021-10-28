package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza Application! ")
	fmt.Println("Please rate our pizza between 1 and 5.")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // without `_` we get "cannot assign 2 values to 1 variables"

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// The input is  actually trailed by `\n` like "4\n" - thus we need TrimSpace

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Final Rating == ", numRating+1)
	}
}
