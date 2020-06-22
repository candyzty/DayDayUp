package main

import (
	"fmt"
	"os"
)

// 定义一个学生的结构体

type persion struct {
	id    int
	name  string
	grade int
	score int
}

func newPerson(id int, name string, grade int, score int) *persion {
	return &persion{
		id,
		name,
		grade,
		score,
	}
}

type class struct {
	Map map[int]*persion
}

func (c *class) showperson() {
	var (
		id int
	)
	//fmt.Println("Check to see if the student exists")
	//fmt.Println("````````````````````````````" )
	//fmt.Print("Please enter student ID:")
	fmt.Scan(&id)
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "ID", "姓名", "年龄", "分数")

	//sortID := make([]int, 0)
	//if sortID != nil{
	//	fmt.Println("input value error!")
	//}
	s := c.Map[id]
	fmt.Println(s.id, s.name, s.grade, s.score)
	//for k := range c.Map {
	//	fmt.Println(sortID, k)
	//}
	//sort.Ints(sortID)
	//for _,k := range sortID{
	//	s := c.Map[k]
	//	fmt.Printf("\t%d\t%s\t%d\t%d\n", s.id, s.name, s.grade, s.score)
	//}
}

func (c *class) addstudent() {
	var (
		id    int
		name  string
		grade int
		score int
	)
	//c := class{}
	fmt.Println("Add student")
	fmt.Print("Input student id: ")
	_, err := fmt.Scan(&id)
	fmt.Print("Input student name: ")
	_, err = fmt.Scan(&name)
	fmt.Print("Input student grade: ")
	_, err = fmt.Scan(&grade)
	fmt.Print("Input student score: ")
	_, err = fmt.Scan(&score)
	a := newPerson(id, name, grade, score)
	//fmt.Println(a.id)
	if err != nil {
		fmt.Println("保存出错！")
	}
	_, isSave := c.Map[id]
	if isSave {
		fmt.Println("学生ID已存在！")
		return
	}
	c.Map[id] = a
	s := c.Map[id]
	fmt.Println(s.id, s.name, s.grade, s.score)

	fmt.Println("Add   student to complete")
}

func (c *class) updatestudent() {
	fmt.Print("update   data")
}

func (c *class) deletestudent() {
	fmt.Print("delete   data")
}
func main() {
	for {
		var choice int
		c := &class{}
		c.Map = make(map[int]*persion, 50)
		//c := class{}
		//a := newPerson(id,name,grade, int8(score))
		fmt.Println("~~~~~ Welcome to the student system ~~~~！")
		fmt.Println(`
1、查看学生信息 2、添加学生信息 3、更新学生信息 4、删除学生信息 5、退出系统
	`)
		fmt.Println("Please enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			c.showperson()
		case 2:
			c.addstudent()
		case 3:
			c.updatestudent()
		case 4:
			c.deletestudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Print("choice  error!")
		}
	}
}
