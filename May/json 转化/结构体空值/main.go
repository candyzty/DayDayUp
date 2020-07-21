package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	//Profile
	//想要在嵌套的结构体为空值时，忽略该字段，仅添加omitempty是不够的：
	//还需要使用嵌套的结构体指针：
	*Profile `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"website"`
	Slogan  string `json:"slogan"`
}

func main() {
	u1 := User{
		Name:  "七米",
		Hobby: []string{"足球", "篮球"},
	}
	b, _ := json.Marshal(u1)
	fmt.Printf("%s\n", b)

}
