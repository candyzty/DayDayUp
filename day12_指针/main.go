package main

import "fmt"

//1. 获取变量的地址
//func main() {
//	var a int = 10
//	fmt.Printf("变量的地址: %x\n", &a)
//}

//2.声明
//func main() {
//	var a int = 20
//	var ip *int
//	ip = &a
//	fmt.Printf("a变量的地址是：%x\n", &a)
//	fmt.Printf("ip 变量的存储地址是：%x\n", ip)
//	fmt.Printf("*ip变量的值：%d\n", *ip)
//}

type name int8
type first struct {
	a int
	b bool
	name
}

//func main() {
//	a := new(first)
//	a.a = 1
//	a.name = 11
//	fmt.Println(a.b, a.a, a.name)
//}
//func main() {
//	var a = first{1, false, 2}
//	var b *first = &a
//	fmt.Println(a.b, a.a, a.name, &a, b.a, &b, (*b).a)
//}

//3.  获取指针的值
//func main() {
//	b := 255
//	a := &b
//	fmt.Println("address of b is", a)
//	fmt.Println("value of b is", *a)
//}

//4.  操作指针改变变量的数值
//func main() {
//	b := 255
//	a := &b
//	fmt.Println("address of b is", a)
//	fmt.Println("value of b is", *a)
//	*a ++
//	fmt.Println("new value of b is", b)
//}

//5. 使用指针传递函数的参数
//func change(x *int)  {
//	*x = 55
//}
//func main() {
//	a := 58
//	fmt.Println("value of a before function call is", a)
//	b := &a
//	change(b)
//	fmt.Println("value of a after function call is", a)
//}

//func modify(arr *[3]int)  {
//	//(*arr)[0] = 100
//	arr[0] = 99
//}
//func main() {
//	a := [3]int{1, 3, 5}
//	modify(&a)
//	fmt.Println(a)
//}

const MAX int = 3

func main() {
	a := [3]int{1, 2, 3}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
