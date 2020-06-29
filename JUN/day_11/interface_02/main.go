package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}

//  A  里面继承了 B  和C  接口，A接口必须把B和C  test03都实现了，才可以使用，缺一不可
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("test01")
}
func (stu Stu) test02() {
	fmt.Println("test02")
}
func (stu Stu) test03() {
	fmt.Println("test03")
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()
	a.test02()
	a.test03()
}
