package main

import "github.com/gin-gonic/gin"

func webHook(c *gin.Context) {

}
func main() {
	r := gin.Default()
	r.POST("/webhook", webHook)
	r.Run(":8080")
}
