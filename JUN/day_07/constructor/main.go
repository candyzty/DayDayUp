package main

import "fmt"

//  构造函数
//  结构体是值类型，赋值的时候就是拷贝

type person struct {
	name string
	age  int
}

// 构造函数： 约定成使用new开头
// 返回person变量
// 构造函数返回的是结构体还是结构体指针
// 当结构体比较大的时候，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}
func main() {
	p1 := newPerson("Tom", 32)
	P2 := newPerson("Jack", 43)
	fmt.Printf("Type:%T value:%v\n", p1, p1)
	fmt.Printf("Type:%T value:%v\n", P2, P2)
	//fmt.Println(P2)
	d1 := newDog("Tibetan Mastiff")
	fmt.Println(d1)
}
