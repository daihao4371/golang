package main

import "fmt"

func main() {
	S1 := "big"
	// 强类型转换
	byteS1 := []byte(S1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) // pig

	S2 := "白萝卜"
	runS2 := []rune(S2)
	runS2[0] = '红'
	fmt.Println(string(runS2)) // 红萝卜
}
