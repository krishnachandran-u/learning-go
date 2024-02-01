package main

import (
    "fmt"
)

func main() {
	//standard for loop
    for i := 0; i < 10; i++ {
        fmt.Println(i, " ")
    }
    fmt.Println()

    //range based loop
    nums := []int{10, 20, 30, 40, 50}

    for index, value := range nums {
        fmt.Println(index, " - ", value)
    }
    fmt.Println()

    //infinite loop
    x := 0
    for {
        fmt.Println(x, " ")
        if x == 7 {
            break
        }
        x++
    }
}
