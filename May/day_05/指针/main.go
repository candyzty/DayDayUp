package main

import "fmt"

// go 语言只有值传递，没有引用传递

// 值传递
func funcValRef(a int) int {
	a = 1000
	return a
}

// 值传递
func funcVal(a *int) {
	*a = 1000
	fmt.Println("指针：", *a)
}
func main() {
	var p int = 100
	fmt.Println(funcValRef(p))
	fmt.Println(p)
	funcVal(&p)
	fmt.Println(p)
}
