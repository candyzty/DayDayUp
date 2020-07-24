package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("loginJson", func(c *gin.Context) {
		var json Login
		// 将request的body中的数据，自动按照json格式解析到结构体

		if err := c.ShouldBindJSON(&json); err != nil {
			// gin.H 封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//	 判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"Status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Status": "200"})
	})
	// 指定服务端口
	r.Run(":9000")

}
