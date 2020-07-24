package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义中间件

func MiddleWare() gin.HandlerFunc {
	t := time.Now()
	return func(c *gin.Context) {
		fmt.Println("中间件开始执行了!")
		// 设置变量到Context的key中，可以通过Get()获取
		c.Set("request", "redraw")
		// 执行中间件
		c.Next()
		statusCode := c.Writer.Status()
		fmt.Println("status:", statusCode)
		t2 := time.Since(t)
		fmt.Println("结束时间:", t2)
	}
}
func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	{
		// 同步
		r.GET("/middleware", func(c *gin.Context) {
			req, _ := c.Get("request")
			//fmt.Println(req)
			c.JSON(200, gin.H{"request": req})
			return
		})
		//	根路由后面定义的是局部中间件
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) {
			req, _ := c.Get("request")
			//fmt.Println(req)
			c.JSON(200, gin.H{"request": req})
			return
		})
	}
	r.Run(":8000")
}
