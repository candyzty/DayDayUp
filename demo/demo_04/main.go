package main

import (
	"DayDayUp/demo/demo_04/controllers"
	"DayDayUp/demo/demo_04/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userV1 := router.Group("/api/user")
	{
		//userV1.GET("/", fetchAllTod)
		userV1.POST("/cre", controllers.CreateUser)
		userV1.PUT("/:id", controllers.UpdateUser)
		userV1.DELETE("/:id", controllers.DeleteUser)
		userV1.POST("/id", controllers.LoginUser)
	}
	db.Init()
	//defer db.Close()
	router.Run(":8001")
}
