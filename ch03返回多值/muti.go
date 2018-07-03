package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

var h, w string = "hello", "world"

func main() {
	a, b := swap(h, w)
	fmt.Println(a, b)
}
