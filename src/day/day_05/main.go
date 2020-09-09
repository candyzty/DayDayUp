package main

import "fmt"

func main() {
	//a := image{name: "张三", age: 20}
	//b := image{}
	//a.printImage()
	//b.age = 20
	//b.name="张三"
	//var c1 *image = &a
	//var c2 *image = &b
	//fmt.Printf("变量a的地址:%X,变量b的地址:%x\n", *c1, *c2)
	//if a == b {
	//	fmt.Println(true);
	//}
	a := new(name)
	a.name = "张三"
	meath := resultMeath(a)
	var i = &meath
	fmt.Println(*i)

	var b = course{"zhang"}
	resultMeath(b)
	//fmt.Println(a.buildName())
	//a.printName()
	//var service nameService = new(name)
	//service.printName()

}

func resultMeath(service nameService) nameService {
	service.printName()
	return service.buildName()
}

type name struct {
	name string
}

type course struct {
	name string
}

func (e course) printName() {
	fmt.Println(e)
}

func (e course) buildName() name {
	n := name{e.name}
	return n
}

type nameService interface {
	printName()
	buildName() name
}

func (e name) printName() {
	fmt.Println(e.name)
}
func (e name) buildName() name {
	return e
}

type image struct {
	name string
	age  int
}

/**
方法。即结构体的方法
*/
func (e image) printImage() {
	fmt.Printf("用户的名字：{%v},年龄：{%v}\n", e.name, e.age)
}
