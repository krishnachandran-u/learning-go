package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Println("Enter a string: ")

	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You entered: ", input)
}
