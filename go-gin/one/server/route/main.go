package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 创建路由
func main() {
	r := gin.Default()

	//  路由组 1，处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	//  路由组 2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	r.Run()
}
func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "tom")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
