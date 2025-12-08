package main

import (
	"log"
)

func main() {
	IFProcess(4)
}

func IFProcess(n int) {
	if n > 5 {
		log.Println("x is greater than 5")
	} else {
		log.Println("x is not greater than 5")
	}
}
