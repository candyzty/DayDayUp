package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int64  `json:"score"`
	Int   int64  `json:"int"`
}

func main() {
	stu1 := student{
		Name:  "沙河",
		Score: 90,
		Int:   19,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.String(), t.Kind(), t.NumField())
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name:%s  index:%d type:%v json tag:%v\n", filed.Name, filed.Index, filed.Type, filed.Tag)
	}
	// 通过字段名称获取结构体字段信息
	if scorefield, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scorefield.Name, scorefield.Index, scorefield.Type, scorefield.Tag)
	}
}
