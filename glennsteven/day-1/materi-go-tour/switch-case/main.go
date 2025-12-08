package main

import "log"

func main() {
	log.Println(SwitchCaseProcess(5))
}

func SwitchCaseProcess(n int) string {
	switch n % 2 {
	case 0:
		return "Genap"
	case 1:
		return "Ganjil"
	default:
		return "Unknown"
	}
}
