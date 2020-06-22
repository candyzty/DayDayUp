package main

import (
	"fmt"
	"os"
)

var srm studentMgr

func showMenu() {
	//fmt.Println()
	fmt.Println("~~~~welcome sms!~~~~")
	fmt.Print(`
	1、查看学生
	2、添加学生
	3、修改学生
	4、删除学生
	5、退出程序
	`)
}
func main() {
	srm = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号:")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			srm.showStudents()
		case 2:
			sys := srm.addStudent()
			fmt.Println(sys)
		case 3:
			srm.putStudent()
		case 4:
			srm.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("序号输入错误！")
		}
		//defer fmt.Print("欢迎下次光临！")
	}
}
