package main

import (
	"github.com/betareduced/gobootcamp/glennsteven/homework-2/internal/types"
)

func main() {
	x, y := 42, 2701
	types.PointerTypes(x, y)

	types.TotalWallet(types.Wallet{Balance: 100}) // Value receiver mengubah COPY, pointer receiver mengubah ASLINYA
	types.BasicArray()

	nums := []int{1, 2, 3, 4}
	types.SliceArray(nums)

	types.SliceUsingMake()

	types.ExampleMaps()
}
