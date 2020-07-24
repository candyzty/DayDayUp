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
	r.POST("loginForm", func(c *gin.Context) {
		var form Login

		// Bind()默认解析并绑定form 格式
		// 根据请求头中 的content-type自动推断

		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"Status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Status": "200"})
	})
	// 指定服务端口
	r.Run(":9000")

}
