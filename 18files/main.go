package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang!")

	var content string
	fmt.Print("Please enter input to be written in file - ")
	fmt.Scan(&content)

	file, err := os.Create("./mygofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("length is - ", length)
	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	// first field is []byte
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Byte data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is - ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
