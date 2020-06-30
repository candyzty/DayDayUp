package main

import (
	"fmt"
	"os"
)

var balance float64
var hcs bulling

func main() {
	//hcs = studentMgr{
	//	allStudent: make(map[int64]bulling, 100),
	//}
	hcs := studentMgr{
		allStudent: make(map[string]bulling, 100),
	}
	number := int(0)
	for {
		fmt.Print("-----家庭收支账本软件-----")
		fmt.Print(`
	1、查看收支明细
	2、收入账单记录
	3、支出账单记录
	4、 退出
	`)
		//fmt.Print(&number)

		fmt.Print("请用户输入选择项:")
		fmt.Scan(&number)
		switch number {
		case 1:
			//fmt.Println("欢迎查看收支明细")
			hcs.LookBill()
		case 2:
			//fmt.Println("欢迎查看收入账单")
			hcs.INcome()
		case 3:
			//fmt.Println("欢迎查看支出账单")
			hcs.Expenditure()
		case 4:
			n := ""
			//fmt.Print("是否真的决定要离开(y/n):")
			fmt.Scan(&n)
			if n == "y" {
				os.Exit(0)
			} else {
				break
			}
		default:
			fmt.Print("输入错误！")
		}
	}
}
