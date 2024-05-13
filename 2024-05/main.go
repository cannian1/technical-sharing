package main

import (
	"fmt"
)

func main() {
	left := 2
	right := 10
	fmt.Println(left, right, (right-left)>>1, (right-left)>>1+left)
	// 2 10 4 6
}
