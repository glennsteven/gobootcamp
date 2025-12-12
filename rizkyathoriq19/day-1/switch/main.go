package main

import "fmt"

func main() {
	weather := "rainy"

	switch weather{
	case "rainy":
		fmt.Println("Today is rainy")
	case "sunny":
		fmt.Println("Today is sunny")
	case "cloudy":
		fmt.Println("Today is cloudy")
	default:
		fmt.Println("Today is windy")
	}
}