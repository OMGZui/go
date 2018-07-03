package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[1:4]
	printSlice(s)

	s = s[:2]
	printSlice(s)

	s = s[1:]
	printSlice(s)

	a := make([]int, 5)
	printSliceBy("a", a)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceBy(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
