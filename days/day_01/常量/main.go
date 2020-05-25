package main

import "fmt"

// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
//
// 不曾使用的常量，在编译的时候，是不会报错的
//
// 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值
// 如果中断iota自增，则必须显式恢复。且后续自增值按行序递增
//
// 自增默认是int类型，可以自行进行显示指定类型
//
// 数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址
const (
	a = iota
	b
	c = iota
)

func main() {
	fmt.Println(a, b, c)

	const name = "edx"
	const (
		x = 16
		y
		s = "asssss"
		z
	)
	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", z, z)

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // 独立值，iota += 1
		e        // "ha"   iota += 1
		f = 100  // iota +=1
		g        // 100  iota +=1   g默认取f的值
		h = iota // 7,恢复计数  这里 继续沿用iota的值 因为上边打断了所以要显式恢复
		i        // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
