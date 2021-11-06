package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi Date and Time!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2021, time.March, 10, 12, 00, 0, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 15:04 Monday"))

	formattedDate := time.Date(2021, time.March, 10, 12, 00, 0, 0, time.UTC).Format("01-02-2006 15:04 Monday")
	fmt.Println(formattedDate)
}
