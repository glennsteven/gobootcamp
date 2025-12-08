package main

import "log"

func main() {
	ForLoopProcess(5)
}

func ForLoopProcess(n int) {
	for i := 0; i < n; i++ {
		log.Println("loop version1:", i)
	}

	i := 0
	for i < n {
		log.Println("loop version2:", i)
		i++
	}
}
