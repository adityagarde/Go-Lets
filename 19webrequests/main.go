package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/adityagarde/Go-Lets"

func main() {
	fmt.Println("Web Requsts!")

	response, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Reponse Type is %T\n", response) // *http.Response

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
