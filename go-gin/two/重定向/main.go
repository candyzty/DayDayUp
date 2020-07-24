package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/redirect", func(c *gin.Context) {
		//	重定向
		//	支持内部和外部重定向
		c.Redirect(301, "http://www.baidu.com/")
	})
	r.Run(":9003")
}
