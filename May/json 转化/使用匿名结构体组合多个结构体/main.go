package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Content string
}
type Image struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func anon() {
	c1 := Comment{
		"永远不要高估自己",
	}
	i1 := Image{
		"暂涨吗",
		"https://www.baidu.com",
	}
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{&c1, &i1})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	jsonStr := `{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}`
	var (
		c2 Comment
		i2 Image
	)
	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)
}
func main() {
	anon()
}
