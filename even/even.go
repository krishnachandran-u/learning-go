package main

import (
	"fmt"
    "os"
)

func main() {
	var num int

	_, err := fmt.Scanln(&num)


    if err != nil {
        fmt.Println("error")
        os.Exit(1)
    }

	if num % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
