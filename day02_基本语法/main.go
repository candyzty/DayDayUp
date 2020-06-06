package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}
func main() {
	////第一种 ：一行声明一个变量
	//var name string = "Python编程时光"
	//var rate float32 = 0.89
	//fmt.Println(name,rate)

	////第二种：多个变量一起声明
	//var (
	//	name string
	//	age int
	//	gender string
	//)
	//fmt.Println(name,age,gender)

	////第三种：声明和初始化一个变量
	//name:="Python编程时光"
	//fmt.Println(name)

	////第四种：声明和初始化多个变量
	//name, age := "wangbm", 28
	//fmt.Println(name,age)

	//第五种：new 函数声明一个指针变量
	//ptr := new(int)
	//fmt.Println("ptr address: ", ptr)
	//fmt.Println("ptr value: ", *ptr)  // * 后面接指针变量，表示从内存地址中取出值

	//匿名变量
	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
}
