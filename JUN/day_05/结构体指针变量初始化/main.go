package main

import "fmt"

// 定义结构体类型

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 顺序初始化，每个成员必须初始化
	var s1 *Student = &Student{
		1,
		"tom",
		'm',
		19,
		"北京",
	}
	fmt.Println(s1)

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := &Student{name: "Mike", addr: "bj"}
	fmt.Println(s2)
}
