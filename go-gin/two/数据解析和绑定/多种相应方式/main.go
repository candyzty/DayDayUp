package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	// 创建路由
	r := gin.Default()
	// 1.json
	r.GET("someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": "200"})
		return
	})
	// 2. 结构体相应
	r.GET("someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "root login succeed"
		msg.Number = 200
		c.JSON(200, msg)
	})
	//	3.XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abd"})
	})
	// 4.yaml
	r.GET("/someYaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "hangman"})
	})

	// 5.protobuf 格式
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
	r.Run(":8000")
}
