package db

import (
	"DayDayUp/demo/demo_04/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:mysql123@tcp(192.168.20.61:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移数据库结构
	db.AutoMigrate(&model.UserRegistration{})
	//db.AutoMigrate(&model.UserRegistration{})
}

//func Close()  {
//	db.Close()
//}
func Create1(str ...interface{}) {
	db.Create(str)
	//db.NewRecord(str)
}
