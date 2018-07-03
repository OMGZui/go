package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	i, j := 1, 2
	fmt.Println(add(i, j))
}
