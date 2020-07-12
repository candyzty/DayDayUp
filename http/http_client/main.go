package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http  client

func main() {
	resp, err := http.Get("http://127.0.0.1:80/xxxx/")
	if err != nil {
		fmt.Print(err)
	}
	//fmt.Println(resp.Body)
	//fmt.Println(resp.Body.Close())
	b, error1 := ioutil.ReadAll(resp.Body)
	if error1 != nil {
		fmt.Println("read resp.body failed", error1)
		return
	}
	fmt.Println(string(b))
}
