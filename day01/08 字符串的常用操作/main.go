package main

import (
	"fmt"
	"strings"
)

func main() {
	// 求长度
	H := "hello world"
	fmt.Println(len(H)) // 11

	// 拼接字符串
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Printf("%s\n", result) // Hello World

	// 判断是否包含
	fmt.Println(strings.Contains(str1, "h")) // false

	// 前缀/后缀判断
	fmt.Println(strings.HasPrefix(str1, "H")) // true
	fmt.Println(strings.HasSuffix(str1, "H")) // false

	// 子串出现的位置
	/*
		strings.Index() 函数用于查找字符串中某个子串第一次出现的位置，
		strings.LastIndex() 函数用于查找字符串中某个子串最后一次出现的位置。
	*/
	// 查找子串第一次出现的位置
	str := "Hello world hello"
	substr := "hello"
	firtsIndex := strings.Index(str, substr)
	fmt.Printf("%s 第一次出现的位置是：%d\n", substr, firtsIndex) // 12

	// 查找子串最后一次出现的位置
	lastIndex := strings.LastIndex(str, substr)
	fmt.Printf("%s, 最后一次出现的位置是：%d\n", substr, lastIndex) // 12

	// join操作
	/*
	 创建一个字符串切片
	 使用 strings.Join() 函数连接字符串切片中的元素
	*/
	a := []string{"apple", "banana", "orange"}
	ret2 := strings.Join(a, ", ") // 使用逗号和空格作为分割符号
	fmt.Println(ret2)             // apple, banana, orange

}
