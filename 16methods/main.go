package main

import "fmt"

func main() {
	fmt.Println("Structs in Go!")

	user01 := User{"Aditya Garde", "test@test.com", true, 25}
	fmt.Println(user01)

	user01.GetStatus()
	user01.NewEMail()

	// Old value is only printed and the local object is changed not the original.
	fmt.Println("Original Email is - ", user01.Email)

}

// First letter capital for Method name OR Variable name or Struct name => It is public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Functions used for Structs are `Methods`
func (user User) GetStatus() {
	fmt.Println("Is user active ? -", user.Status)
}

func (user User) NewEMail() {
	user.Email = "testing@test.com"
	fmt.Println("New user email is - ", user.Email)
}
