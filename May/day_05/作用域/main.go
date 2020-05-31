package main

import "fmt"

//变量作用域：
//
//局部变量：函数内定义的变量（作用域只在函数体内，参数和返回值变量也是局部变量）
//全局变量：函数外定义的变量（全局变量可以在整个包甚至外部包（被导出后）使用）
//形式参数：函数定义中的变量（作为函数的局部变量来使用）

// 全局变量
var g int = 100

func sum(a, b int) int {
	var c int
	s := a + b + c + g
	return s
}

func main() {
	// 局部变量
	var a, b, c int
	a = 10
	b = 20
	c = 15
	fmt.Println(a + b + c)
	fmt.Println(sum(a, b))
}
