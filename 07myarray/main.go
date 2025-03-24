package main

import (
	"fmt"
)

func main() {

	fmt.Println("Lets learn Arrys in GO")

	var authors [5]string

	authors[0] = "A"
	authors[1] = "B"
	authors[2] = "C"
	authors[3] = "D"
	authors[4] = "E"

	fmt.Println("Authors List : ", authors)
	fmt.Println("Authors List : ", len(authors))

	var reader = [3]string{"Nik", "vik", "lik"}
	fmt.Println("Reader List is", reader)

}
