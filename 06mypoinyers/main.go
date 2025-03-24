package main

import "fmt"

func main() {
	fmt.Println("Lets learn Pointers")

	// var ptr *int
	// fmt.Println("Vale of Pointer is", ptr)

	mynumber := 23
	var ptr = &mynumber
	fmt.Println("Value of Pointer is", ptr)
	fmt.Println("Value of Pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is", mynumber)

}
