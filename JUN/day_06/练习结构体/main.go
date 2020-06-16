package main

import (
	"fmt"
	"log"
	"sort"
)

type Student struct {
	ID    int
	Name  string
	Age   int8
	Score int8
}

type Class struct {
	Map map[int]*Student
}

func (c *Class) AddStudent() {
	var id int
	var name string
	var age int8
	var score int8
	fmt.Println("输入id: ")
	_, err := fmt.Scan(&id)
	fmt.Print("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入年龄: ")
	_, err = fmt.Scan(&age)
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
	student := &Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	c.Map[id] = student
	//fmt.Println("保存成功！")
	log.Println("保存成功！")
}
func (c *Class) ShowStudent() {
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "ID", "姓名", "年龄", "分数")
	sortID := make([]int, 0)
	for k := range c.Map {
		sortID = append(sortID, k)
	}
	sort.Ints(sortID)
	for _, k := range sortID {
		s := c.Map[k]
		fmt.Printf("\t%d\t%s\t%d\t%d\n", s.ID, s.Name, s.Age, s.Score)
	}
}

// 删除学生

func (c *Class) DeleteStudent() {
	fmt.Println("输入要删除学生的ID:")
	var id int
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("err 报错了！")
	}
	_, isSave := c.Map[id]
	if !isSave {
		fmt.Println("要删除的ID不存在！")
		return
	}
	delete(c.Map, id)
	fmt.Println("删除成功！")
}
func (c *Class) UpdateStudent() {
	var id int
	fmt.Println("输入修改的学生ID:")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("err  报错！")
	}
	_, isSave := c.Map[id]
	if !isSave {
		fmt.Println("要修改的ID不存在！")
		return
	}
	var name string
	var age int8
	var score int8
	fmt.Println("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入年龄: ")
	_, err = fmt.Scan(&age)
	fmt.Print("输入分数: ")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("保存出错！")
	}
	student := &Student{
		ID:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	c.Map[id] = student
	fmt.Println("修改成功！")
}
func main() {
	c := &Class{}
	c.Map = make(map[int]*Student, 50)
	for {
		fmt.Println("要执行的操作：")
		fmt.Println("1. 添加  2.查看  3.删除  4.修改")
		var do int8
		_, err := fmt.Scan(&do)
		if err != nil {
			fmt.Println("输入有误！")
		}
		switch do {
		case 1:
			c.AddStudent()
		case 2:
			c.ShowStudent()
		case 3:
			c.DeleteStudent()
		case 4:
			c.UpdateStudent()
		default:
			fmt.Println("输入有误！")
		}
	}
}
