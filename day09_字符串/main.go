package main

import "fmt"

//1.显示
//func main() {
//	name := "Hello World"
//	fmt.Println(name)
//}

//2. 访问字符串中的单个字节
func main() {
	name := "Hello World"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%d ", name[i])
	}
	fmt.Println()
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
}
