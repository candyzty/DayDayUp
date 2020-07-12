package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//str := "hello world!"
	b, err := ioutil.ReadFile("F:\\code\\go\\src\\DayDayUp\\http\\http_service\\readme.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("暂时无内容 %v\n", err)))
	}
	w.Write([]byte(b))
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("Ok"))
}
func main() {
	http.HandleFunc("/hello/", hello)
	http.HandleFunc("/xxxx/", f2)
	http.ListenAndServe("0.0.0.0:80", nil)
}
