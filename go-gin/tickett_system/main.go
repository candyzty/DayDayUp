package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func httpStatus() string {
	//var   ma = make(map[string]string)
	url := "http://192.168.8.62:10810/actuator/health"
	//for k,v := range ma {
	//	fmt.Println(k, v)
	http, err := http.Get(url)
	body, err := ioutil.ReadAll(http.Body)
	if err != nil {
		return "url 报错！"
	}
	return string(body)
}

func main() {
	r := gin.Default()
	a := httpStatus()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"meaafe":  a,
			"status":  200,
			//a := httpStatus()
			//fmt.Println(a)
		})
	})
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
