package ServiceModule

import "fmt"

// 定义结构体
type Customer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Phone  int64  `json:"phone"`
	Email  string `json:"email"`
}

func NewCustomer(name string, gender string, age int, phone int64, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

type CustomerInfo struct {
	//customers []*Customer
	CustomerNub int
	Customers   []*Customer
	//customers   interface{}

	//customerNub int
	//customers   int

	//customerNub interface{}
	//customers   interface{}
}

func (c *CustomerInfo) List() {
	fmt.Println("查看用户")
	fmt.Println("--------------------客户列表--------------------")
	fmt.Println("编号\t  姓名\t  性别\t  年龄\t  电话号码\t\t   邮箱\t ")
	for _, customer := range c.Customers {
		fmt.Printf("%+v\n", customer)
	}
}

//type customers struct {
//
//}

func (c *CustomerInfo) Add() {
	// c := &CustomerInfo{
	//	1,
	//	{
	//
	//		gender string
	//		age    int
	//		phone  int64
	//		email  string
	//	},
	// }
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
	c.CustomerNub++
	// ClientInfo := &NewCustomer{name,gender,age,phone,email}
	stu := &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
	c.Customers = append(c.Customers, stu)
}
func (c *CustomerInfo) Update() {
	fmt.Println("修改用户")
}
func (c *CustomerInfo) Delete() {
	fmt.Println("查删除用户")
}

func main() {

}
