package main

import (
	"fmt"
	"time"
)

const FormatDate = `2006-01-02 15:04:05`

func main() {
	now := time.Now()
	fmt.Println(now.Format(FormatDate))
}
