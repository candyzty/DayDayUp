package main

import (
	"DayDayUp/JUN/crs/service"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("----欢迎来到客户管理系统----")
	for {
		fmt.Print("----欢迎来到客户管理系统----")
		fmt.Print(`
	1、查看用户
	2、增加用户
	3、修改用户
	4、删除用户
	5、退出系统
	`)
		number := int(0)
		fmt.Print("请选择(1-5):")
		fmt.Scan(&number)
		switch number {
		case 1:
			service.NewCustomerService().List()
			//fmt.Println(list)
		case 2:
			AddInfo := service.Add()
			fmt.Println(AddInfo)
		case 3:
			fmt.Println("修改用户")
		case 4:
			fmt.Println("删除用户")
		case 5:
			n := ""
			fmt.Print("是否真的决定要离开(y/n):")
			fmt.Scan(&n)
			if n == "y" {
				os.Exit(0)
			} else {
				break
			}
		default:
			fmt.Println("参数错误，请继续输入！")
		}
	}
}
