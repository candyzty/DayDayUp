package main

import (
	"fmt"
)

// 定义结构体类型

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	// 定义一个结构体普通变量
	var s Student
	s.id = 1
	s.name = "mike"
	s.sex = 'w'
	s.age = 19
	s.addr = "hz"
	fmt.Println(s)

	// 指针有合法指向后，才操作成员
	s1 := new(Student)
	s1.id = 2
	s1.name = "Tom"
	s1.sex = 'w'
	s1.age = 32
	fmt.Println(s1)
}
