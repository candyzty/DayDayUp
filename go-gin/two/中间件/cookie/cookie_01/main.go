package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "Not set"
			fmt.Println("cookie  error:", err)
			//	 设置cookie
			c.SetCookie("key_cookie", "value_cookie", 120, "/", "localhost", false, true)
		}
		fmt.Println(cookie)
	})
	r.Run(":9000")
}
