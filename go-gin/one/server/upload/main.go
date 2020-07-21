package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"net/http"
)

func main() {
	r := gin.Default()
	// 限制表单上传的大小 8MB,默认值32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 表单取文件
		//file, err := c.FormFile("file")
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}

		//log.Print(file.Filename)
		// 上传文件的路径
		// 上传单个文件
		//fielder := "F:/code/go/src/DayDayUp/go-gin/one/server/imageFile/"
		// 上传多个文件

		//fielder := "F:/code/go/src/DayDayUp/go-gin/one/server/imageFiles/"
		//path := fielder + file.Filename

		files := form.File["files"]

		for _, file := range files {
			fielder := "F:/code/go/src/DayDayUp/go-gin/one/server/imageFiles/"
			path := fielder + file.Filename
			c.SaveUploadedFile(file, path)
			//fmt.Println(a)

			log.Print(file.Filename)
		}
		// 打印xinx
		c.String(200, fmt.Sprintf("%v upload image!", len(files)))
	})
	r.Run(":8000")
}
