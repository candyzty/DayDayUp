package main

import (
	"DayDayUp/demo/demo_02/controllers"
	"DayDayUp/demo/demo_02/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()
	v1 := r.Group("/v1")
	//v1.Use()
	{
		v1.GET("/people", controllers.GetPerson)
		v1.POST("/people/:id", controllers.CreatePerson)
		v1.PUT("/people/:id", controllers.UpdatePerson)
		v1.DELETE("/people/:id", controllers.Delete)
	}
	r.Run(":8000")
	defer db.CloseDB()
}
