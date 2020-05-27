package main

import (
	"fmt"
)

// 运算符
// GO语言中没有前置得++和--，++a、--a是错误得

func operatorDemo() {
	var a int = 20
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("a + b = %d\n", c)
	c = a - b
	fmt.Printf("a - b = %d\n", c)

	c = a * b
	fmt.Printf("a * b = %d\n", c)
	c = a % b
	fmt.Printf("a %% b = %d\n", c)
	c = a / b
	fmt.Printf("a / b = %d\n", c)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

}
func relational() {
	var a = 21
	var b = 10
	if a == b {
		fmt.Printf("a==b\n")
	} else {
		fmt.Printf("a != b\n")
	}
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}
}

func logicDemo() {
	//	真假判断
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("a和b都是真")
	} else {
		fmt.Println("a和b有假")
	}
	if a || b {
		fmt.Println("a和b中有一个为真")
	} else {
		fmt.Println("a和b都为false")
	}

	if !a {
		fmt.Println("a为假")
	} else {
		fmt.Println("a为真")
	}

}
func byteComputeDemo() {
	//var a int = 60
	var a uint = 60
	var b uint = 13
	var c uint
	c = a & b
	fmt.Printf("a&b=%d\n", c)

	c = a | b
	// a = 111000 0
	// b = 001101 1
	fmt.Printf("a|b=%d\n", c)
	c = a ^ b
	// a = 111000 0
	// b = 001101 1
	fmt.Printf("a^b=%d\n", c)
	c = a << 2
	// a << n 左移n位=乘以2得n次方
	fmt.Printf("a<<2=%d\n", c)
	c = b >> 2
	// b  >> n 右移n位=除以2得n次方
	fmt.Printf("b>>2=%d\n", c)
}
func main() {
	//operatorDemo()b
	//relational()
	//logicDemo()
	byteComputeDemo()
}
