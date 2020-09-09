package main

import "fmt"

func main() {
	teacher := teacher{name: "王亮"}
	student := student{name: "张三", age: 20, teacher: teacher}
	autor := autor{lastName: "王大妈", fristName: "王三四", student: student}
	autor.printName()
}

/**
结构体1
*/
type autor struct {
	fristName string
	lastName  string
	student
}

func (e autor) printName() {
	fmt.Printf("第一个名字是：%s,第二个名字是：%s\n", e.fristName, e.lastName)
	fmt.Printf("autor 下的学生姓名：%s,学生年龄：%d\n", e.student.name, e.student.age)
	fmt.Printf("该学生负责的老师是：%s\n", e.student.teacher.name)
}

/**
结构体2
*/
type student struct {
	name      string
	age       int8
	className string
	teacher
}

/**
结构体3
*/
type teacher struct {
	name string
}
