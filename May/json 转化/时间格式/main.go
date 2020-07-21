package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	CreateTime time.Time `json:"create_time"`
}

func timeFieldDemo() {
	p1 := Post{CreateTime: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time": "2020-07-21  18:34:21"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
	// 错误信息
	//	json.Unmarshal failed, err:parsing time ""2020-07-21  18:34:21"" as ""2006-01-02T15:04:05Z07:00"": cannot parse "
	//	18:34:21"" as "T"
}

func main() {
	timeFieldDemo()
}
