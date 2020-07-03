package ServiceModule

import (
	"bufio"
	"fmt"
	"os"
)

// 定义结构体
type Customer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Phone  int64  `json:"phone"`
	Email  string `json:"email"`
}

//func NewCustomer(id int, name string, gender string, age int, phone int64, email string) Customer {
//	return Customer{
//		Id:     id+1,
//		Name:   name,
//		Gender: gender,
//		Age:    age,
//		Phone:  phone,
//		Email:  email,
//	}
//}

type CustomerInfo struct {
	//CustomerNub int
	Customers []*Customer
}

// 定义输出的格式
func (c *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t %v\t %v\t\t %v\t\t %v\t %v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}

func (c *CustomerInfo) List() {
	fmt.Println("查看用户")
	fmt.Println("--------------------客户列表--------------------")
	fmt.Println("编号\t  姓名\t  性别\t  年龄\t  电话号码\t\t   邮箱\t ")
	//fmt.Println(len(c.Customers))
	if 0 > len(c.Customers) {
		fmt.Println("没有数据，请先添加数据进来")
	} else {
		for i := 0; i < len(c.Customers); i++ {
			info := c.Customers[i].GetInfo()
			fmt.Println(info)
		}
	}
	fmt.Println()
}

// 添加客户信息
func (c *CustomerInfo) Add() {
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
	stu := &Customer{
		Id:     +1,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
	c.Customers = append(c.Customers, stu)

}

// 修改客户信息
func (c *CustomerInfo) Update() {
	for i := 0; i < len(c.Customers); i++ {
		info := c.Customers[i].GetInfo()
		fmt.Println(info)
	}
	fmt.Println(len(c.Customers))

	a := c.Customers[1]
	fmt.Println(a)
	fmt.Println("--------------------修改客户--------------------")
	var (
		Name   string
		Gender string
		Age    int
		Phone  int64
		Email  string
	)
	fmt.Printf("姓名：%v 修改为", a.Name, Name)
	fmt.Scan(&Name)
	//fmt.Printf("姓名：%v 修改为",c.Name,Name)
	//fmt.Scan(&Name)
	fmt.Printf("性别：%v 修改为", a.Gender, Gender)
	fmt.Scan(&Gender)
	fmt.Printf("年龄：%v 修改为", a.Age, Age)
	fmt.Scan(&Age)
	fmt.Printf("号码：%v 修改为", a.Phone, Phone)
	fmt.Scan(&Phone)
	fmt.Printf("邮箱：%v 修改为", a.Email, Email)
	fmt.Scan(&Email)

	for i := 0; i < len(c.Customers); i++ {
		info := c.Customers[i].GetInfo()
		fmt.Println(info)
	}
	fmt.Println(len(c.Customers))
}

// 删除客户信息
func (c *CustomerInfo) Delete() {
	fmt.Println("查删除用户")
}

// 终端输入，如果中间有空格，使用scan只能接收第一个参数

// 使用bufio.NewReader
func useBufio() {
	s := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容为:%s\n", s)
}
