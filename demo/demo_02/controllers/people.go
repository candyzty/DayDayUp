/**
* @Author: redgr
* @Description:
* @File:  people
* @Version: 1.0.0
* @Date: 2020/7/26 22:25
 */

package controllers

import (
	"DayDayUp/demo/demo_02/db"
	"DayDayUp/demo/demo_02/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ error

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var person models.Person
	var getDB = db.GetDB()
	d := getDB.Where("id=?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{
		"id #" + id: "deleted"})
}
func UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	var getDB = db.GetDB()
	if err := getDB.Where("id = ?", id).First(&person).Error; err != nil {
		fmt.Println("更新失败", err)
	}
	//c.BindJSON(&person)
	getDB.Save(&person)
	c.JSON(http.StatusOK, person)
}
func CreatePerson(c *gin.Context) {
	var person models.Person
	var getDB = db.GetDB()

	//c.BindJSON(&person)
	getDB.Create(&person)
	c.JSON(http.StatusOK, person)
}
func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)
	var person models.Person
	var getDB = db.GetDB()
	if err := getDB.Where("id = ?", id).First(&person).Error; err != nil {
		fmt.Println("获取参数失败", err)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
