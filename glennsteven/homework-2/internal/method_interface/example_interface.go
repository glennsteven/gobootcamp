package method_interface

import "fmt"

type Count struct {
	ValueA, ValueB int
}

func (c Count) Sum() int {
	return c.ValueA + c.ValueB
}

func (c *Count) ChangeValue(ValueC int) {
	c.ValueA = c.ValueA * ValueC
	c.ValueB = c.ValueB * ValueC
}

// Resolver Interfaces
// Menggunakan interface, berarti harus mengimplementasikan suatu method sesuai dengan aturan yang ada di interface
type Resolver interface {
	Sum() int
}
type SumResolver struct {
	A, B int
}

func (s *SumResolver) Sum() int {
	return s.A + s.B
}

func TypeAssertionInterface(val interface{}) {
	s, ok := val.(string)
	fmt.Println("Assertion string:", s, ok)

	i, ok := val.(int)
	fmt.Println("Assertion integer:", i, ok)

	f, ok := val.(float64)
	fmt.Println("Assertion float:", f, ok)
}

// IPAddr Stringers -> Exercise
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		ip[0], ip[1], ip[2], ip[3],
	)
}
