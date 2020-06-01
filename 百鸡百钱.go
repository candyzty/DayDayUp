package main

import "fmt"

// 公鸡5文钱一只，母鸡3文钱一只，小鸡3只一文钱， 用100文钱买一百只鸡,其中公鸡，母鸡，小鸡都必须要有，问公鸡，母鸡，小鸡要买多少只刚好凑足100文钱？
func main() {
	for x := 0; x < 20; x++ {
		for y := 0; y < 33; y++ {
			var z = 100 - x - y
			if z%3 == 0 && 5*x+3*y+z/3 == 100 && x != 0 && y != 0 {
				fmt.Printf("公鸡:%v只%v文,母鸡:%v只%v文,小鸡:%v只%v文\n", x, x*5, y, y*3, z, z/3)
			}
		}
	}
}
