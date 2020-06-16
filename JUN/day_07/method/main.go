package main

import "fmt"

// 方法

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name,
	}
}

// 方法是作用域特定类型的函数
// 接收者表示的是调用该方法的具体的类型变量，用类型名首字母小写表示
func (d dog) wang() {
	//d.name = "小黑"
	fmt.Printf("%s: 汪汪汪！\n", d.name)
}
func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}

func (p *person) guonian() {
	p.age++
}

func (p *person) dream() string {
	dream := "好好学习"
	return dream
}
func main() {
	d1 := newDog("xiaobai")
	d1.wang()
	p1 := newPerson("xiaobai", 32)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	a := p1.dream()
	fmt.Println(a)
}
