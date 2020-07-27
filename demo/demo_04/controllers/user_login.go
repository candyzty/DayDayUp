package controllers

import (
	"DayDayUp/demo/demo_04/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//var db *gorm.DB

type UserRegistration struct {
	//User     User `json:"user"`
	//gorm.Model
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

func CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	phone, _ := strconv.Atoi(c.PostForm("phone"))
	fmt.Println(email)
	fmt.Println(phone)
	userInfo := UserRegistration{
		//User: &model.User{
		//	Username: user,
		//	Password: password,
		//},
		Phone: phone,
		Email: email,
	}
	fmt.Println(userInfo)
	db.Create1(userInfo)
	//db.Create1(&userInfo)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user create success!",
	})
	//defer db.Close()
}

func LoginUser(c *gin.Context) {

}
func UpdateUser(c *gin.Context) {

}
func DeleteUser(c *gin.Context) {

}
