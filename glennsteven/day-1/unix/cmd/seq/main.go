package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		_, err := fmt.Fprintln(os.Stderr, "usage: seq N")
		if err != nil {
			log.Println("error process usage:", err)
			return
		}
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, "invalid number:", os.Args[1])
		if err != nil {
			log.Println("error process check argument:", err)
			return
		}
		os.Exit(1)
	}

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
