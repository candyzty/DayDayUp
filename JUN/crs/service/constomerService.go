package service

import (
	"DayDayUp/JUN/crs/module"
	"fmt"
	//"fmt"
)

type CustomerService struct {
	customers []module.Customer
	// 声明一个字段，表示当前还切片里面有多少个客户
	customerNub int
}

func NewCustomerService() *CustomerService {
	constomerservice := &CustomerService{}
	constomerservice.customerNub = 1
	cusromer := module.NewCustomer(1, "Tom", "男", 23, 18888888888, "tom@163.com")
	//fmt.Printf("%T\n",cusromer)
	constomerservice.customers = append(CustomerService{}.customers, cusromer)

	fmt.Printf("constomerservice.customers :%v\n", &constomerservice.customers)
	return constomerservice
}

// 返回客户切片
func (this *CustomerService) List() {
	customers := this.customers
	fmt.Println(customers)
	//customers[i].GetInfo()
	fmt.Println("--------------------客户列表--------------------")
	fmt.Println("编号\t  姓名\t  性别\t  年龄\t  电话号码\t\t   邮箱\t ")
	for i := 0; i < len(customers); i++ {
		//customers := this.customers
		info := customers[i].GetInfo()
		fmt.Println(info)
	}
	fmt.Println("------------------客户列表完成------------------")
}
func (this *CustomerService) Add() []module.Customer {
	constomerservice := &CustomerService{}
	fmt.Println("--------------------添加客户--------------------")
	var (
		name   string
		gender string
		age    int
		phone  int64
		email  string
	)

	fmt.Print("请输入姓名:")
	fmt.Scan(&name)
	fmt.Print("请输入性别:")
	fmt.Scan(&gender)
	fmt.Print("请输入年龄:")
	fmt.Scan(&age)
	fmt.Print("请输入手机号:")
	fmt.Scan(&phone)
	fmt.Print("请输入邮箱:")
	fmt.Scan(&email)
	this.customerNub++
	//this.customerNub =  this.customerNub
	//constomerservice.customerNub = id + 1
	// 去ID
	cusromer := module.NewCustomer(this.customerNub, name, gender, age, phone, email)
	constomerservice.customers = append(this.customers, cusromer)
	//this.customers = append(CustomerService{}.customers,cusromer)
	//this.customers= append(this.customers,cusromer)

	fmt.Println("--------------------添加完成--------------------")
	//fmt.Println(this.customers)
	return constomerservice.customers
}

//func Add()  *CustomerService {
//	constomerservice := &CustomerService{}
//	var(
//		name string
//		gender string
//		age  int
//		phone  int64
//		email  string
//	)
//
//	fmt.Print("请输入姓名:")
//	fmt.Scan(&name)
//	fmt.Print("请输入性别:")
//	fmt.Scan(&gender)
//	fmt.Print("请输入年龄:")
//	fmt.Scan(&age)
//	fmt.Print("请输入手机号:")
//	fmt.Scan(&phone)
//	fmt.Print("请输入邮箱:")
//	fmt.Scan(&email)
//
//	constomerservice.customerNub ++
//	cusromer := module.NewCustomer(constomerservice.customerNub,name,gender,age,phone,email)
//	//fmt.Printf("%T\n",cusromer)
//	constomerservice.customers = append(CustomerService{}.customers,cusromer)
//	fmt.Printf( "constomerservice.customers :%v\n", &constomerservice.customers)
//	return constomerservice
//}
