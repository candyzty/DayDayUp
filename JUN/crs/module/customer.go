package module

import (
	"fmt"
)

// 声明一个客户信息结构体

type Customer struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Phone  int64  `json:"phone"`
	Email  string `json:"email"`
}

func NewCustomer(id int, name string, gender string, age int, phone int64, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 不带id的
func NewCustomer2(name string, gender string, age int, phone int64, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t %v\t %v\t\t %v\t\t %v\t %v", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return info
}
