package main

import (
	"fmt"
	"os"
	"sort"
)

// 定义一个学生的结构体

type persion struct {
	id     int
	name   string
	grade  int
	score  int
}

func newPerson(id  int,name string, grade int, score int)  *persion  {
	return &persion{
		id,
		name,
		grade,
		score,
	}
}

type class struct {
	Map  map[int]*persion

}

func (c  *class)  showperson() {
	var(
		id int
	)
	fmt.Println("Check to see if the student exists")
	fmt.Print("Please enter student ID")
	fmt.Scan(&id)
	sortID := make([]int, 0)
	//student := c.Map[id]
	for k := range c.Map {
		sortID = append(sortID, k)
	}
	sort.Ints(sortID)
}

func  (c  *class)  addstudent() {
	var(
		id int
		name string
		grade int
		score int
	)
	//c := class{}
	fmt.Println("Add student")
	fmt.Print("Input student id: ")
	fmt.Scan(&id)
	fmt.Print("Input student name: ")
	fmt.Scan(&name)
	fmt.Print("Input student grade: ")
	fmt.Scan(&grade)
	fmt.Print("Input student score: ")
	fmt.Scan(&score)
	//student := &persion{
	//	id,
	//	name,
	//	grade,
	//	score,
	//}
	//a := &persion{
	//	id,
	//	name,
	//	grade,
	//	score,
	//}
	a := newPerson(id,name, grade, score)
	fmt.Println(a.id)
	//a.id = id
	//a.name = name
	//a.grade = grade
	//a.score = int8(score)
	c.Map[id] = a
	fmt.Print("Add   student to complete")
}
func main() {
	var choice int
	//var(
	//	id int
	//	name string
	//	grade string
	//	score int
	//)
	c := &class{}
	c.Map = make(map[int]*persion, 50)
	//c := class{}
	//a := newPerson(id,name,grade, int8(score))
	fmt.Print("~~~~~ Welcome to the student system ~~~~！")
	fmt.Println(`
1、查看学生信息 2、添加学生信息 3、更新学生信息 4、删除学生信息 5、退出系统
	`)
	fmt.Print("Please enter your choice: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		c.showperson()
	case 2:
		c.addstudent()
	case 3:
		//c.updatestudent()
	case 4:
		//c.deletestudent()
	case 5:
		os.Exit(0)
	default:
		fmt.Print("choice  error!")
	}
}

