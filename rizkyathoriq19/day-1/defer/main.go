package main

import "fmt"

func main() {
    defer fmt.Println("Last Program")

    fmt.Println("First Program")
}
