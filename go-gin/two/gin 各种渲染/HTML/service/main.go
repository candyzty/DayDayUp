package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 指定文件加载
	//r.LoadHTMLFiles("templates/index.tmpl")
	// 全局加载
	r.LoadHTMLGlob("E:/GO学习小分队/src/github.com/edifierx666/DayDayUp/go-gin/two/gin 各种渲染/HTML/templates/*")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染
		// 根据json将title替换
		c.HTML(200, "index.tmpl", gin.H{
			"title": "我的标题",
		})
	})
	r.Run(":9002")
}
