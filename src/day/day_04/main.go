package main

import (
	"fmt"
)

/**
指针的学习
指针是存储另一个变量的内存地址的变量
获取指针地址在指针变量前加&的方式
取地址符 ： &
 */
func main() {
	test2()
}
/**
声明指针
 */
func stment(){
	var ip *int
	a := 10
	ip = &a
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", ip)
	fmt.Printf("%d\n", *ip)
}

/*
结构体
 */
type first struct {
	name string
	age int
	b bool
}
func test1(){
	a := first{"zhang",21,true}
	//a.age = 21
	//a.name = "张三"
	fmt.Printf("%+v\n", a)
	var b *first = &a
	fmt.Printf("地址：%x\n", &b)
	(*b).age = 20
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)

}

func test2(){
	x := 10
	y := 20
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}

/**
交换a和b的指针地址
 */
func swap(x *int, y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}