package main

import "fmt"

func main() {
	fmt.Println("Structs in Go!")

	// no inheritance and related concepts (like super, parent-child classes etc.) in Golang.

	user01 := User{"Aditya Garde", "test@test.com", true, 25}
	fmt.Println(user01)                           // {Aditya Garde test@test.com true 25}
	fmt.Printf("user01 details == %+v\n", user01) // {Name:Aditya Garde Email:test@test.com Status:true Age:25}
	fmt.Printf("user01's Name is %v and email is %v.\n", user01.Name, user01.Email)

}

// First letter capital for Method name OR Variable name or Struct name => It is public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
