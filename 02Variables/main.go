package main

import "fmt"

func main() {
	var shambho string = "Nikhil"
	fmt.Println(shambho)
	fmt.Printf("type of veriable: %T \n", shambho)

	var Isloggedin bool = true
	fmt.Println(Isloggedin)
	fmt.Printf("type of veriable: %T \n", Isloggedin)

	var smallval int = 255
	fmt.Println(smallval)
	fmt.Printf("type of veriable: %T \n", smallval)

	var usmall uint = 256
	fmt.Println(usmall)
	fmt.Printf("type of veriable: %T \n", usmall)

	var Smallf float64 = 256.8747347
	fmt.Println(Smallf)
	fmt.Printf("type of veriable: %T \n", Smallf)

	// implicit type
	var Shambo = "Nik"
	fmt.Println(Shambo)

	// no var style

	learn := "Learn golang"
	fmt.Println(learn)
}
