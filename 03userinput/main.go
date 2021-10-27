package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user Input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza : ")

	// comma ok OR comma err syntax - substitution for try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)      // 5
	fmt.Printf("Type of this rating is %T", input) // string!
}
