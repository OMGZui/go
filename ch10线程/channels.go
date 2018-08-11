package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

/*
➜  ch10线程 git:(master) ✗ go run channels.go
6 -5 12 17
➜  ch10线程 git:(master) ✗ go run channels.go
6 17 12 -5
➜  ch10线程 git:(master) ✗ go run channels.go
6 -5 12 17
➜  ch10线程 git:(master) ✗ go run channels.go
6 17 -5 12
➜  ch10线程 git:(master) ✗ go run channels.go
17 -5 12 6
➜  ch10线程 git:(master) ✗ go run channels.go
6 -5 17 12
*/
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go sum(s, c)
	go sum([]int{1, 2, 3}, c)
	x, y, z, r := <-c, <-c, <-c, <-c
	fmt.Println(x, y, z, r)
}
