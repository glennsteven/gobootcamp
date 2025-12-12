package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args [0] = program name
	// [1], [2] = argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: seq <n>")
        return
    }

	// ASCII to Integer
    n, _ := strconv.Atoi(os.Args[1])

    for i := 1; i <= n; i++ {
        fmt.Println(i)
    }
}
