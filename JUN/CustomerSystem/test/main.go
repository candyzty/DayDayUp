package ServiceModule

import "fmt"

// 定义结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  int64
	Email  string
}

type CustomerInfo struct {
	customerNub int
	customers   Customer
}

var data = []CustomerInfo{
	CustomerInfo{1, Customer{1, "2", "3", 4, 5, "6"}},
	CustomerInfo{2, Customer{1, "2", "3", 4, 5, "6"}},
}

func (c *CustomerInfo) List() {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
	fmt.Printf("%v", data)

}

func (c *CustomerInfo) Add() {
	fmt.Println("添加用户")
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
	fmt.Println("增加前", c.customerNub)
	c.customerNub++
	fmt.Println("增加后", c.customerNub)
	//ClientInfo := &NewCustomer{name,gender,age,phone,email}
	stu := Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
	c.customers = stu

	data = append(data, *c)
}
func (c *CustomerInfo) Update() {
	fmt.Println("修改用户")
}
func (c *CustomerInfo) Delete() {
	fmt.Println("查删除用户")
}
