package main

import (
	"flag"
	"fmt"
)

func main() {
	// 创建一歌标志位参数
	name := flag.String("name", "李四", "请输入名字")
	age := flag.Int("age", 999, "请输入年龄")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}
