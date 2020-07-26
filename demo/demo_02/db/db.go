/**
* @Author: redgr
* @Description:
* @File:  db
* @Version: 1.0.0
* @Date: 2020/7/26 22:15
 */

package db

import (
	"DayDayUp/demo/demo_02/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Init() {
	db, _ = gorm.Open("mysql", "root:mysql123@tcp(192.168.20.61:3306)/go_test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.Person{})
}

// GetDB
func GetDB() *gorm.DB {
	return db
}
func CloseDB() {
	db.Close()
}
