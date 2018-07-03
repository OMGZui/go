package main

import "fmt"

type V struct {
	X int
	Y int
}

func main() {
	fmt.Println(V{1, 2})
	v := V{1, 2}
	v.X = 4
	fmt.Println(v)
	// 通过指针修改结构体的值
	p := &v
	p.X = 1e9
	fmt.Printf("%T %v\n", v, v)
}
