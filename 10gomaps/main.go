package main

import "fmt"

func main() {
	fmt.Println("Maps in Go!")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["GO"] = "Golang"
	languages["PY"] = "Python"
	languages["JV"] = "Java"
	fmt.Println(languages)

	fmt.Println("JS = ", languages["JS"])

	delete(languages, "PY")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("Language == %v\n", value)
	}
}
