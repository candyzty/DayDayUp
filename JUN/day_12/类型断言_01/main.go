package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	var a interface{}
	var point Point = Point{10, 12}
	a = point
	var b Point
	b = a.(Point)
	//b = a
	fmt.Println(b)

	//类型断言的案例
	var x interface{}
	var z float32 = 1.1
	x = z
	//y := x.(float32)
	//fmt.Printf("y 的类型是%T y 的值为:%v", y,y)

	//  待检测
	if y, ok := x.(float32); ok {
		fmt.Println("类型错误")
		fmt.Printf("y 的类型是%T y 的值为:%v", y, y)
	} else {
		fmt.Println("convert fail!")
	}
	fmt.Println("继续执行！")
}
