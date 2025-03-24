package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok // Err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for feedback, ", input)
	fmt.Printf("type of this %T", input)

}
