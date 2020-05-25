package main

import "log"

// 变量必须先定义才能使用
// go语言是静态语言，要求变量的类型和赋值的类型必须一致。
// 变量名不能冲突。(同一个作用于域内不能冲突)
// 简短定义方式，左边的变量名至少有一个是新的
// 简短定义方式，不能定义全局变量。
// 变量的零值。也叫默认值。
// 变量定义了就要使用，否则无法通过编译。
func main() {
	// 显示声明
	var name string
	name = "edx"
	// 推导式
	var nam2 = "edx"
	nam3 := "edx" // 只能在函数内
	var name4, name5 = "a1", "a2"
	var name12 string
	log.Println(name, nam2, nam3, name4, name5, name12 == "")
}
