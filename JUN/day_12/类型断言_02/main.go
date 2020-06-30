package main

import "fmt"

// 定义结构体类型
type Student struct {
	Name string
}
type sss struct {
}

// 编程一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for i, v := range items {

		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数为bool类型，值为:%v\n", i, v)
		case int:
			fmt.Printf("第%v个参数为int类型，值为:%v\n", i, v)
		case string:
			fmt.Printf("第%v个参数为string类型，值为:%v\n", i, v)
		case float32:
			fmt.Printf("第%v个参数为float32类型，值为:%v\n", i, v)
		case float64:
			fmt.Printf("第%v个参数为float64类型，值为:%v\n", i, v)
		case Student:
			fmt.Printf("第%v个参数为Student类型，值为:%v\n", i, v)
		case sss:
			fmt.Printf("第%v个参数为struct类型，值为:%v\n", i, v)
		default:
			fmt.Printf("第%v个参数为不确定类型，值为:%v\n", i, v)
		}
	}
}
func main() {
	n1 := float32(1.1)
	n2 := float64(2.3)
	n3 := 43
	name := "tom"
	address := "北京"
	b := 's'
	stu := Student{
		Name: "Tom",
	}
	s := sss{}
	TypeJudge(n1, n2, n3, name, address, b, stu, s)
}
