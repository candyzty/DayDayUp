package main

import (
	"encoding/json"
	"fmt"
)

//  在tag中添加omitempty忽略空值
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
}

//   结构体里面的参数没有值，会输出默认值
//   str:{"name":"气你","email":"","hobby":null}
func main() {
	u1 := User{
		Name: "气你",
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal,err:%s\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

}
