package main

import "fmt"

// 定义接口

type AInterface interface {
	//name string  // 接口中不能有任何变量
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println(stu.Name)

}

type Benefice interface {
	hello()
}
type Monster struct {
}

//func (m Monster) Say() {
//	panic("implement me")
//}

func (m Monster) hello() {
	fmt.Println("Monster hello()")
}
func (m Monster) Say() {
	fmt.Println("Monster say")
}
func main() {
	var stu Stu
	stu.Name = "周晨"
	var a AInterface = stu
	a.Say()

	var monster Monster
	var b1 AInterface = monster
	var b2 Benefice = monster
	b1.Say()
	b2.hello()

}
