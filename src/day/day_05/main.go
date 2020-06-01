package main

import "fmt"

func main(){
	a := image{name: "张三", age: 20}
	b := image{}
	a.printImage()
	b.age = 20
	b.name="张三"
	var c1 *image = &a
	var c2 *image = &b
	fmt.Printf("变量a的地址:%X,变量b的地址:%x\n", *c1, *c2)
	if a == b {
		fmt.Println(true);
	}

}
type image struct {
	name string
	age int
}
/**
方法。即结构体的方法
 */
func (e image) printImage() {
	fmt.Printf("用户的名字：{%v},年龄：{%v}\n", e.name,e.age)
}
