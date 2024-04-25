package main

import "fmt"

/*
// 全局变量
var m = 100

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
}

*/

/*
匿名变量
*/

func foo() (int, string) {
	return 10, "summer"
}

func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
