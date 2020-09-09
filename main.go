package main

import (
	"fmt"
	"log"
)

func main() {
	//rocketmq.Init()
	var class class = class{Map: map[int]*persion{}}

	exit(class)

}
func exit(class class) bool {
	var a string
	fmt.Println("选择模式：add(新增用户)，info(班级信息)")
	fmt.Scanln(&a)
	if a == "add" {
		class.addstudent()
	} else if a == "info" {
		class.printInfo()
	} else if a == "exit" {
		return true
	}
	exit(class)
	return false
}

type persion struct {
	id int
	name string
	grade string
	score int8
}


type class struct{
	Map map[int]*persion
}

func (c *class) addstudent() {
	var id int
	var name string
	var grade string
	var score int8
	fmt.Println("输入id: ")
	_, err := fmt.Scan(&id)
	fmt.Print("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入班级: ")
	_, err = fmt.Scan(&grade)
	fmt.Print("输入分数: ")
	_, err = fmt.Scan(&score)
	if err != nil {
		log.Println("保存出错！")
	}
	_, isSave := c.Map[id]
	if isSave {
		log.Println("学生ID已存在！")
		return
	}
	student := &persion{
		id,
		name,
		grade,
		score,
	}
	log.Println(student)
	c.Map[id] = student
	//fmt.Println("保存成功！")
	log.Println("保存成功！")

}

func (c *class) printInfo() {
	for i := range c.Map {
		log.Println(*c.Map[i])
	}
}