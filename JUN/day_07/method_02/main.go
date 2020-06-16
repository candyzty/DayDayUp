package main

import "fmt"

// 给自定义 int 类型加方法

type i int

func (a i) hello() int {
	return int(a)
}
func main() {
	m := i(1000)
	fmt.Println(m.hello())
}
