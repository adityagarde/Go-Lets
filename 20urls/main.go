package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://github.com/adityagarde?tab=overview&from=2021-12-01&to=2021-12-31"

func main() {
	fmt.Println("Playing with URLS in Go!")
	fmt.Println(myurl)

	// parsing the URL
	result, err := url.Parse(myurl)
	checkNilError(err)

	fmt.Println("result.Scheme is ", result.Scheme)
	fmt.Println("result.Host is ", result.Host)
	fmt.Println("result.Path is", result.Path)
	fmt.Println("result.Port() is", result.Port())
	fmt.Println("result.RawQuery is", result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of queryParams is %T\n", queryParams) //url.Values - key, value pairs

	fmt.Println(queryParams["tab"])

	for key, value := range queryParams {
		fmt.Printf("%v : %v\n", key, value)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "github.com",
		Path:     "adityagarde",
		RawQuery: "tab=repositories",
	}

	fmt.Println(partsOfUrl.String())
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
