package main

import "fmt"

// myVariable := "Some Value"; // Walrus Symbol (:=) Not Allowed outside methods
var myVariable = "Some Value"

const LoginToken string = "Sherlock Holmes" // variable name starts with Capital letter => Public

func main() {
	var username string = "Aditya Garde"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n ", username)

	var isBoolean bool = true
	fmt.Println(isBoolean)
	fmt.Printf("Variable is of type : %T \n ", isBoolean)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n ", smallValue)

	var intValue int = 256
	fmt.Println(intValue)
	fmt.Printf("Variable is of type : %T \n ", intValue)

	var float32Value float32 = 255.787809123423678
	fmt.Println(float32Value) //255.78781
	fmt.Printf("Variable is of type : %T \n ", float32Value)

	var float64Value float64 = 255.787809123423678
	fmt.Println(float64Value) //255.78780912342367
	fmt.Printf("Variable is of type : %T \n ", float64Value)

	//implicit type
	var someString = "Aditya Garde"
	fmt.Println(someString)
	fmt.Printf("Variable is of type : %T \n ", someString)

	// no var style
	someNumber := 347209
	fmt.Println(someNumber)
	fmt.Printf("Variable is of type : %T \n ", someNumber)

	fmt.Println(myVariable)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n ", LoginToken)

}
