package main

import (
	"fmt"
	"math"
)

// 常量 枚举学习
// 常量的关键字： const
// 常量的值不能进行改变
// 枚举是一种特殊的常量，可以通过iota快速设置连续的值
// 类型定义与类型别名
// 类型定义：基于int创建的一个新类型，主要提高代码可读性
// 类型别名：基于int创建的一个别名，和原类型完全一样，主要用于包兼容
// Go 语言中常量一般不用大写，大写表示public,可导出的，有特殊含义

// 定义常量
func constDemo() {
	const s string = "hello"
	const a, b = 3, 4
	fmt.Println(s, a, b)

	var c int
	c = int(math.Sqrt(float64(c)))
	fmt.Println(c)

	const (
		s1  = "golang"
		max = 10
	)
	fmt.Println(s1, max)
}

// 枚举类型 iota
func enumDemo() {
	const (
		Monday = 1 + iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}

func typeDefAndAlias() {
	type MyInt1 int   // 基于int 定义的一个新的类型,这个用的多 (别名)
	type MyInt2 = int // MyInt和int 是完全一样的 type alias

	var i int = 1
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)

}
func main() {
	constDemo()
	enumDemo()
	typeDefAndAlias()
}
