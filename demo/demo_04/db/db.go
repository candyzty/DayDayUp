package db

import (
	"DayDayUp/demo/demo_04/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 初始化mysql 信息
func Init() {
	var err error
	DB, err = gorm.Open("mysql", "root:mysql123@tcp(192.168.20.61:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移数据库结构
	DB.AutoMigrate(&model.UserRegistration{})
	//db.AutoMigrate(&model.UserRegistration{})
}
