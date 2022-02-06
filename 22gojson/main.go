package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling JSON in Go!")

	EncodeJSON()

	DecodeJSON()
}

func EncodeJSON() {

	myCourses := []course{
		{"Apache Kafka Mastery", 499, "Udemy", "kafka123", []string{"Java", "Backend", "Event-sourcing"}},
		{"Reactive Microservices with Webflux", 499, "Udemy", "webflux123", []string{"Java", "Backend", "Spring"}},
		{"Introduction to Go", 0, "Youtube", "go123", []string{"Go", "New"}},
		{"AWS for Java Developers", 499, "Udemy", "aws123", []string{"Java", "AWS", "Cloud"}},
		{"Refactoring your Resume", 0, "Youtube", "resume123", nil},
	}

	// packaging the data as JSON
	// finalJSON, err := json.Marshal(myCourses) // No Indenteation
	finalJSON, err := json.MarshalIndent(myCourses, "", "\t") // prefix and Indent added
	checkNilError(err)

	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDummyData := []byte(`
	{
		"courseName": "Reactive Microservices with Webflux",
		"Price": 499,
		"Platform": "Udemy",
		"tags": ["Java", "Backend", "Spring"]
	}
	`)

	var myCourse course

	isValid := json.Valid(jsonDummyData)

	if isValid {
		json.Unmarshal(jsonDummyData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Invalid JSON received!")
	}

	var myDataMap map[string]interface{}
	json.Unmarshal(jsonDummyData, &myDataMap)
	fmt.Printf("%#v\n", myDataMap)

	for key, value := range myDataMap {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, value, value)
	}
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
