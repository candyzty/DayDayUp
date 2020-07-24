package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("程序用时:", since)
}

func shopIndex(c *gin.Context) {
	time.Sleep(5 * time.Second)

}
func shopHome(*gin.Context) {
	time.Sleep(7 * time.Second)
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(MyTime)
	shoppingGroup := r.Group("/shopping")
	{
		// 同步
		shoppingGroup.GET("/index", shopIndex)

		//	根路由后面定义的是局部中间件
		shoppingGroup.GET("/home", shopHome)
	}
	r.Run(":8000")
}
