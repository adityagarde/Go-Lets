package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Requests in Go!")

	const getURL = "http://localhost:8000/get"
	const postURL = "http://localhost:8000/post"
	const postformURL = "http://localhost:8000/postform"

	PerformGetRequest(getURL)
	PerformPostJsonRequest(postURL)
	PerformPostFormRequest(postformURL)
}

func PerformGetRequest(myURL string) {

	response, err := http.Get(myURL)
	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	// option 1
	fmt.Println(content) // byte format
	fmt.Println(string(content))

	// option 2
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count:", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest(myURL string) {

	// dummy payload
	requestBody := strings.NewReader(`
	{
		"name":"Aditya Garde",
		"age":"25"
	}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)
	checkNilError(err)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostFormRequest(myURL string) {

	// dummy form data
	data := url.Values{}
	data.Add("name", "Aditya Garde")
	data.Add("age", "25")

	response, err := http.PostForm(myURL, data)
	checkNilError(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
