package main

import (
	"fmt"
)

/**
第三天。
函数
函数是Go语言中的一种数据类型，可以作为另一个函数的参数，也可以作为函数的返回值

defer函数
延迟语句，被用于执行一个函数调用，在这个函数之前，延迟语句返回
类似与java的finaaly
*/
func main() {
	day_05()
	//局部变量
	//var param1 int8 = 1
	//param2 := "hello"
	//param3 := "word"
	//name, s := funcName(param1, param2)
	//fmt.Printf("param1:%d,param2:%s",name,s)
	//s2 := test1(param2, param3)
	//fmt.Println(s2);
	//test3(&param1)
	//fmt.Println(param1);
	//test1(param4, param5)
	//fmt.Println(param6);

}

//全局变量.
//全局变量不能使用自动推导类型的方式来定义变量
var (
	param4 string = "xiix"
	param5 string = "haah"
)

var param6 string = "vsvsv"

/**
学习函数的
入参，出参
入参：实参，形参
实参：值传递。Go语言的基本数据类都为值传递
形参：引用传递。即指针传递，将变量的指针地址copy到参数上
.slice,map。可以直接传递。需要改变slice长度时需要取地址传递指针
*/
func funcName(param1 int8, param2 string) (int8, string) {
	return param1, param2
}

//可变参数.
func test1(arg ...string) string {
	fmt.Printf("param:{%v}", arg)
	return "1"
}

//形参
func test3(a *int8) {
	*a++
}

/**
延迟函数
defer
在代码段结束执行之前，一定执行的方法
*/
func day_05() {
	name := "Naveen"
	fmt.Println("Orignal string :" + name)
	fmt.Printf("Reversed string : ")
	for i, _ := range name {
		defer fmt.Printf("%c", name[i])
	}
	fmt.Println("结束啦。哈哈哈哈")
}
