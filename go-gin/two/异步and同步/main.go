package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	// 1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要弄一个副本，不能直接操作
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行:" + copyContext.Request.URL.Path)
		}()
		return
	})
	//	2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行:" + c.Request.URL.Path)
		return
	})
	r.Run(":9004")
}
