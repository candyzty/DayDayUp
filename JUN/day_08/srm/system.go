/**
* @Author: redgr
* @Description:
* @File:  system
* @Version: 1.0.0
* @Date: 2020/6/20 10:25
 */

package main

import "fmt"

type student struct {
	id   int64
	name string
}
type studentMgr struct {
	allStudent map[int64]student
}

func (s studentMgr) estimate(stuId int64) {
	_, err := s.allStudent[stuId]
	if !err {
		fmt.Print("ID不存在")
		return
	} else {
		fmt.Print("操作成功！")
	}
}

func (s studentMgr) showStudents() {
	var choice int
	fmt.Print(`
	1、查看所有学生
	2、查询单个学生
	`)
	fmt.Print("请输入序号:")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		for k, v := range s.allStudent {
			fmt.Println(k, v.name)
		}
	case 2:
		var stuId int64
		fmt.Print("请输入学生ID：")
		fmt.Scan(&stuId)
		//_,ok := s.allStudent[stuId]
		//if !ok {
		//	fmt.Print("ID不存在")
		//	return
		//}
		s.estimate(stuId)
		fmt.Println(s.allStudent[stuId])
	default:
		fmt.Print("查询完成！")
	}
}
func (s studentMgr) addStudent() map[int64]student {
	var (
		stuId   int64
		stuName string
	)
	fmt.Print("请输入学号:")
	fmt.Scan(&stuId)
	fmt.Print("请输入姓名:")
	fmt.Scan(&stuName)
	newStu := student{
		stuId,
		stuName,
	}
	s.allStudent[newStu.id] = newStu
	return s.allStudent

}
func (s studentMgr) putStudent() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号:")
	fmt.Scan(&stuID)
	subObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Print("ID不存在")
		return
	} else {
		fmt.Printf("hello %T, value:%v", subObj, subObj)
	}
	fmt.Printf("你要修改学生的信息如下，学号:%d 姓名:%s \n", subObj.id, subObj.name)
	fmt.Print("请输入姓名:")
	fmt.Scan(&stuName)
	subObj.name = stuName
	s.allStudent[stuID] = subObj
	fmt.Printf(subObj.name)
}
func (s studentMgr) deleteStudent() {
	var (
		stuId int64
	)
	fmt.Print("输入你要删除的学号:")
	fmt.Scan(&stuId)
	//_, ok := s.allStudent[stuID]
	//if !ok {
	//	fmt.Print("学号ID不存在！")
	//	return
	//}
	s.estimate(stuId)
	delete(s.allStudent, stuId)
	//if stuId  {
	//	fmt.Print("操作Id为:", stuId)
	//}else {
	//	fmt.Print(nil)
	//}
}
