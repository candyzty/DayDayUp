package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func JsonToMapDemo() map[string]string {
	file, _ := os.Open("url_status\\url_check\\url.json")
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		content = append(content, tmp[:n]...)
	}

	var mapResult map[string]string
	err := json.Unmarshal([]byte(content), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	return mapResult
}

func urlCheck(url map[string]string) {
	for i, v := range url {
		fmt.Println(i, v)

		_, err := http.Get(v)
		if err != nil {
			//fmt.Println(err.Error("url 不存在"))
			fmt.Println(errors.New("url 连接失败"))
			//fmt.Sprintf()
		}
		//status := response.StatusCode
		//fmt.Println(status)
		//if status != 200{
		//	fmt.Println(errors.New("url 连接失败"))
		//}
		//fmt.Println("连接成功！")
	}
}

func main() {
	urlCheck(JsonToMapDemo())
	//map1 := JsonToMapDemo()
	//fmt.Println(map1)
}
