package main

import "log"

func main() {
	// 显示声明
	var name string
	name = "edx"
	// 推导式
	var nam2 = "edx"
	nam3 := "edx" // 只能在函数内

	log.Println(name, nam2, nam3)
}
