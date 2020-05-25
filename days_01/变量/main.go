package main

// package 定义main 包
import "fmt"

// 全局的变量
var a int    // int 类型默认值为：0，
var b string // string 类型默认值为：空值.
var c, d int = 1, 2
var e, f = 1, "go"

// 定义多个变量
var (
	l = 25
	m = "Becca"
	n = "3.54"
)

// main 函数只能有一个
func main() {
	// 推导类型，只能在函数中定义
	name := "Tom"
	age := 21
	set := 4.5423
	fmt.Println(name, age, set)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d)
	fmt.Println(e, f, true)
	fmt.Printf("l = %d m = %s n = %v", l, m, n)
}
