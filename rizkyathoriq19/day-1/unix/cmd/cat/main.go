package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args [0] = program name
	// [1], [2] = argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: cat <file>")
        return
    }

	// data = []byte
	// err = error
    data, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(data))
}
