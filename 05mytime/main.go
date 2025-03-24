package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time things")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15.05.33 Monday"))

	createdate := time.Date(2025, time.March, 23, 10, 41, 56, 0, time.UTC)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("01-02-2006 Monday"))

}
