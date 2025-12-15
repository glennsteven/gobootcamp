package main

import (
	"fmt"
	"github.com/betareduced/gobootcamp/glennsteven/homework-2/internal/method_interface"
)

func main() {
	count := method_interface.Count{ValueA: 3, ValueB: 8}
	fmt.Println("Simple Method:", count.Sum())

	count1 := method_interface.Count{ValueA: 3, ValueB: 8}
	count1.ChangeValue(3)
	fmt.Println("Pointer Receiver:", count1.Sum())

	var r method_interface.Resolver
	resolver := method_interface.SumResolver{
		A: 10,
		B: 9,
	}

	r = &resolver
	fmt.Println("Interface Resolver:", r.Sum())

	method_interface.TypeAssertionInterface("Hello, World!")
	method_interface.TypeAssertionInterface(1)
	method_interface.TypeAssertionInterface(1.3)

	fmt.Println("Exercise Method Stringers")
	hosts := map[string]method_interface.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
