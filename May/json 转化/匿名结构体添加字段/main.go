package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func anonymousStructDemo() {
	u1 := UserInfo{
		ID:   1231,
		Name: "杭州",
	}
	b, err := json.Marshal(struct {
		*UserInfo
		Token string `json:"token"`
	}{
		&u1,
		"32234234",
	})
	if err != nil {
		fmt.Printf("json.marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%s\n", b)
}
func main() {
	anonymousStructDemo()
}
