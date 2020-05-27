package main

import (
	"fmt"
)

// switch 默认就有break
// 匹配到的 case 后面不需要加 break，相当于默认就有 break；默认情况下 case 匹配成功后就不会执行其他 case，
// 如果我们需要执行后面的 case，
// 可以使用 fallthrough, fallthrough 不会判断下一条 case 的表达式结果是否为 true。
func main() {
	a := 210
	if a < 20 {
		fmt.Printf("a<10")
	} else if a >= 10 {
		fmt.Printf("a=10")
	} else {
		fmt.Printf("a>10")
	}
	fmt.Println("--------------")
	c := 30
	switch c {
	case 20:
		fmt.Println("c=20")
	case 10:
		fmt.Println("a=10")
	case 15:
		fmt.Println("c=15")
	default:
		fmt.Println("c=30")

	}

}
