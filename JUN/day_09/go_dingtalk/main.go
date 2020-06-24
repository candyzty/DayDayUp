package main

import (
	"fmt"
	"net/http"
)

func main() {
	resonne, err := http.Get("http://www.baidu.com")
	resonne.StatusCode = '3'

	if err != nil {
		fmt.Println("url  错误")
	}
	//resonne.StatusCode = 201
	//fmt.Println(s)

	r := resonne.StatusCode
	fmt.Println(r)

}
