package main

import "fmt"

func maxString(s string) int {
	last := make(map[rune]int)
	start := 0
	max := 0

	for i, v := range []rune(s) {
		if lastI, ok := last[v]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > max {
			max = i - start + 1
		}
		last[v] = i
	}
	return max
}

func main() {
	fmt.Println(maxString("abcacbbb"))
	fmt.Println(maxString("bbbb"))
	fmt.Println(maxString("pwwkew"))
	fmt.Println(maxString("我爱你我"))
}
