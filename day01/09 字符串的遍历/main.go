package main

import "fmt"

func main() {
	// 遍历字符串
	s := "深圳彭于晏"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
}
