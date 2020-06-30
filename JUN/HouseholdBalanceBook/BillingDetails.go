package main

import (
	"fmt"
)

//var balance float64

type bulling struct {
	Income        string
	SystemBalance float64
	IncomeBalance float64
	Explain       string
}

type studentMgr struct {
	allStudent map[string]bulling
}

// 查看收支明细
func (bu studentMgr) LookBill() {
	fmt.Println("查看收支明细")
	//for _, i := range bu. {
	//	fmt.Printf( "",bu.Income,bu.SystemBalance,bu.IncomeBalance,bu.Explain)
	//}
	if balance <= 0 {
		fmt.Println("没有收入支出明细，请先进行添加收入")
	} else {
		fmt.Println("收  支\t 账号余额\t 收支余额\t 说  明\t")
		for _, v := range bu.allStudent {

			//fmt.Println("收  支\t 账号余额\t 收支余额\t 说  明\t")t
			//fmt.Print("收支\t%v 账号余额\t%v 收支余额\t%v 说  明\t%v \n",)
			fmt.Println(v.Income, "\t", v.SystemBalance, "\t", v.IncomeBalance, "\t", v.Explain, "\t")
		}
	}
}

// 登记收入

func (bu studentMgr) INcome() {
	fmt.Println("登记收入")
	var (
		money   float64
		explain string
	)
	fmt.Print("本次收入金额:")
	fmt.Scan(&money)
	fmt.Print("本次收入说明:")
	fmt.Scan(&explain)
	balance = money + balance
	newStu := bulling{
		Income:        "收入",
		SystemBalance: balance,
		IncomeBalance: money,
		Explain:       explain,
	}
	bu.allStudent[newStu.Explain] = newStu
}

//  登记支出
func (bu studentMgr) Expenditure() {
	fmt.Println("登记收入")
	var (
		money   float64
		explain string
	)
	fmt.Print("本次支出金额:")
	fmt.Scan(&money)
	fmt.Print("本次支出说明:")
	fmt.Scan(&explain)

	//if float64(money)　> float64(balance) {
	//	fmt.Println("余额不足，请适当消费")
	//	break
	//}
	if money > balance {
		fmt.Println("余额不足请适当消费！")
	}
	balance = balance - money
	newStu := bulling{
		Income:        "支出",
		SystemBalance: balance,
		IncomeBalance: money,
		Explain:       explain,
	}
	bu.allStudent[newStu.Explain] = newStu
}
