package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// 定义数据结构
type Car struct {
	gorm.Model
	Brand   string `json:"brand"`
	Yielded string `json:"yielded"`
}

type Spec struct {
	Car    Car
	Carboy float64 `json:"carboy"`
	Drive  string  `json:"drive"`
}

func Insert() {
	spec := Spec{Car: Car{Brand: "宝马", Yielded: "德国"}, Carboy: 5.96, Drive: "四驱"}
	//spec1 := Spec{aa:Car{Brand:"宝马", Yielded:"德国"}, Carboy: 5.96, Drive: "四驱"}
	//spec1 := Spec{Car:Car{Brand:"宝马", Yielded:"德国"}, Carboy: 5.96, Drive: "四驱"}
	//ss := Spec{Carboy: 33.4,Drive: "前驱"}
	fmt.Println(spec)
	db.Create(&spec)
	//db.Create(&ss)
}

// 初始化mysql 信息
func main() {
	var err error
	db, err = gorm.Open("mysql", "root:mysql123@tcp(192.168.20.61:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移数据库结构
	db.AutoMigrate(&Car{})
	db.AutoMigrate(&Spec{})
	//db.AutoMigrate(&model.UserRegistration{})
	Insert()
	defer db.Close()
}
