package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Plain Old For Loop
	for i := 0; i < len(days); i++ {
		fmt.Print(days[i] + " ")
	}
	fmt.Println()

	// Even in this loop we get the index and not the value and thus we need to paas days[i] and NOT just i.
	for i := range days {
		fmt.Print(days[i] + " ")
	}
	fmt.Println()

	// Similar to for..each loop
	for index, value := range days {
		fmt.Printf("Index is %v and value is %v\n", index, value)
	}

	// By _ we can skip either the index or the actual value.
	for _, value := range days {
		fmt.Printf("Value is %v\n", value)
	}

	// Similar to while loop
	someValue := 1
	for someValue < 10 {
		if someValue == 5 {
			someValue++
			continue
		} else if someValue == 8 {
			goto goloops
		}
		if someValue == 9 {
			break
		}
		fmt.Println("Value here is - ", someValue)
		someValue++
	}

goloops:
	fmt.Println("Label goLoops reached.")
}
